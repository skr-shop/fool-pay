package common

import (
	"github.com/openpeng/fool-pay/notify"
)

type ReqData struct {
	Body           string `json:"body"`
	Subject        string `json:"subject"`
	Openid         string `json:"openid"`
	OrderNo        string `json:"order_no"`
	TimeoutExpress int64  `json:"timeout_express"`
	Amount         string `json:"amount"`
	ReturnParam    string `json:"return_param"`
	// 支付宝公有
	GoodsType int8   `json:"goods_type"`
	StoreId   string `json:"store_id"`

	// 条码支付
	OperatorId    string `json:"operator_id"`
	TerminalId    string `json:"terminal_id"`
	AlipayStoreId string `json:"alipay_store_id"`
	Scene         string `json:"scene"`
	AuthCode      string `json:"auth_code"`

	// Web支付
	QrMod     string `json:"qr_mod"`
	Paymethod string `json:"paymethod"`

	ClientIp  string `json:"client_ip"`
	ProductId string `json:"product_id"`
}

type ResData struct {
	notify.NotifyProcessData
}

type BaseData struct {
	ReqData ReqData
	ResData ResData
}
