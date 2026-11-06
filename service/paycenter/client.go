package paycenter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	httpTimeout   = 15 * time.Second
	maxBodyBytes  = 1 << 20
	channelWechat = "WECHAT"
	payTypeNative = "NATIVE"
)

var httpClient = &http.Client{Timeout: httpTimeout}

type Config struct {
	BaseURL string
	APIKey  string
}

type APIError struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("支付中心请求失败（HTTP %d）", e.HTTPStatus)
}

type CreateNativeOrderRequest struct {
	BusinessOrderNo string
	AmountFen       int64
	Subject         string
	Attach          string
	ExpireMinutes   int
}

type CreateNativeOrderResult struct {
	OrderNo string
	Status  string
	Amount  int64
	CodeURL string
}

type Order struct {
	OrderNo         string `json:"orderNo"`
	BusinessOrderNo string `json:"businessOrderNo"`
	Amount          int64  `json:"amount"`
	Status          string `json:"status"`
}

type envelope struct {
	Code    int             `json:"code"`
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type createOrderBody struct {
	BusinessOrderNo string `json:"businessOrderNo"`
	Amount          int64  `json:"amount"`
	Subject         string `json:"subject"`
	Channel         string `json:"channel"`
	PayType         string `json:"payType"`
	Attach          string `json:"attach,omitempty"`
	ExpireMinutes   int    `json:"expireMinutes,omitempty"`
}

type createOrderData struct {
	OrderNo   string          `json:"orderNo"`
	Status    string          `json:"status"`
	Amount    int64           `json:"amount"`
	PayParams json.RawMessage `json:"payParams"`
}

type nativePayParams struct {
	CodeURL string `json:"codeUrl"`
}

func (c Config) CreateNativeOrder(ctx context.Context, req CreateNativeOrderRequest) (*CreateNativeOrderResult, error) {
	body := createOrderBody{
		BusinessOrderNo: req.BusinessOrderNo,
		Amount:          req.AmountFen,
		Subject:         req.Subject,
		Channel:         channelWechat,
		PayType:         payTypeNative,
		Attach:          req.Attach,
		ExpireMinutes:   req.ExpireMinutes,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	env, err := c.do(ctx, http.MethodPost, "/api/pay/orders", payload)
	if err != nil {
		return nil, err
	}
	var data createOrderData
	if len(env.Data) > 0 && string(env.Data) != "null" {
		if err := json.Unmarshal(env.Data, &data); err != nil {
			return nil, err
		}
	}
	codeURL := ""
	if len(data.PayParams) > 0 && string(data.PayParams) != "null" {
		var params nativePayParams
		if err := json.Unmarshal(data.PayParams, &params); err != nil {
			return nil, err
		}
		codeURL = params.CodeURL
	}
	return &CreateNativeOrderResult{
		OrderNo: data.OrderNo,
		Status:  data.Status,
		Amount:  data.Amount,
		CodeURL: codeURL,
	}, nil
}

func (c Config) QueryByBusinessOrderNo(ctx context.Context, businessOrderNo string) (*Order, error) {
	path := "/api/pay/orders/business/" + url.PathEscape(businessOrderNo)
	env, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var order Order
	if len(env.Data) > 0 && string(env.Data) != "null" {
		if err := json.Unmarshal(env.Data, &order); err != nil {
			return nil, err
		}
	}
	return &order, nil
}

func (c Config) CloseOrder(ctx context.Context, orderNo string) error {
	path := "/api/pay/orders/" + url.PathEscape(orderNo) + "/close"
	_, err := c.do(ctx, http.MethodPost, path, nil)
	return err
}

func (c Config) do(ctx context.Context, method, path string, payload []byte) (*envelope, error) {
	base := strings.TrimRight(strings.TrimSpace(c.BaseURL), "/")
	if base == "" || strings.TrimSpace(c.APIKey) == "" {
		return nil, fmt.Errorf("支付中心未配置")
	}
	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, base+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(io.LimitReader(resp.Body, maxBodyBytes))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		message := "请求过于频繁，请稍后再试"
		var env envelope
		if json.Unmarshal(raw, &env) == nil && env.Message != "" {
			message = env.Message
		}
		return nil, &APIError{HTTPStatus: resp.StatusCode, Code: http.StatusTooManyRequests, Message: message}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("支付中心 HTTP %d", resp.StatusCode)
	}

	var env envelope
	if err := json.Unmarshal(raw, &env); err != nil {
		return nil, err
	}
	if env.Code != http.StatusOK || !env.Success {
		message := env.Message
		if message == "" {
			message = "支付中心请求失败"
		}
		return nil, &APIError{HTTPStatus: resp.StatusCode, Code: env.Code, Message: message}
	}
	return &env, nil
}
