package paycenter

import "github.com/shopspring/decimal"

// YuanToFen 把人民币元四舍五入成整数分。分是支付中心金额的唯一来源。
func YuanToFen(yuan float64) int64 {
	return decimal.NewFromFloat(yuan).Mul(decimal.NewFromInt(100)).Round(0).IntPart()
}

// FenToYuan 由整数分还原元，写入本地订单的 Money。
func FenToYuan(fen int64) float64 {
	value, _ := decimal.NewFromInt(fen).Div(decimal.NewFromInt(100)).Float64()
	return value
}
