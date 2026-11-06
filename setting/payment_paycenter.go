package setting

import "strings"

var PayCenterEnabled = false
var PayCenterBaseURL = ""
var PayCenterApiKey = ""
var PayCenterNotifyKey = ""
var PayCenterProductCode = ""
var PayCenterUnitPrice = 7.3
var PayCenterMinTopUp = 1

// PayCenterReady 开关打开且下单、验签所需配置都齐全。
func PayCenterReady() bool {
	return PayCenterEnabled &&
		strings.TrimSpace(PayCenterBaseURL) != "" &&
		strings.TrimSpace(PayCenterApiKey) != "" &&
		strings.TrimSpace(PayCenterNotifyKey) != "" &&
		strings.TrimSpace(PayCenterProductCode) != ""
}
