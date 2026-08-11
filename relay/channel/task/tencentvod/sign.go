package tencentvod

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const (
	tencentVODHost    = "vod.tencentcloudapi.com"
	tencentVODService = "vod"
	tencentAlgorithm  = "TC3-HMAC-SHA256"
)

// buildAuthorization generates a TC3-HMAC-SHA256 Authorization header value.
// host and service are parameterized so the same function works for both VOD
// (vod.tencentcloudapi.com / vod) and AI3D (ai3d.tencentcloudapi.com / ai3d).
//
//	action  — e.g. "CreateAigcImageTask", "SubmitHunyuanTo3DRapidJob"
//	payload — raw JSON request body
//	host    — e.g. tencentVODHost or tencentAI3DHost
//	service — e.g. tencentVODService or tencentAI3DService
func buildAuthorization(secretID, secretKey, action, payload string, ts int64, host, service string) string {
	t := time.Unix(ts, 0).UTC()
	date := t.Format("2006-01-02")

	// Step 1: canonical request
	payloadHash := sha256hex([]byte(payload))
	canonicalHeaders := fmt.Sprintf("content-type:application/json\nhost:%s\nx-tc-action:%s\n",
		host, strings.ToLower(action))
	signedHeaders := "content-type;host;x-tc-action"

	canonicalRequest := strings.Join([]string{
		"POST",
		"/",
		"",
		canonicalHeaders,
		signedHeaders,
		payloadHash,
	}, "\n")

	// Step 2: string to sign
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	stringToSign := strings.Join([]string{
		tencentAlgorithm,
		fmt.Sprintf("%d", ts),
		credentialScope,
		sha256hex([]byte(canonicalRequest)),
	}, "\n")

	// Step 3: signing key
	secretDate := hmacSHA256([]byte("TC3"+secretKey), []byte(date))
	secretService := hmacSHA256(secretDate, []byte(service))
	secretSigning := hmacSHA256(secretService, []byte("tc3_request"))
	signature := hex.EncodeToString(hmacSHA256(secretSigning, []byte(stringToSign)))

	// Step 4: Authorization header value
	return fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		tencentAlgorithm, secretID, credentialScope, signedHeaders, signature)
}

func sha256hex(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
