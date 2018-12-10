package common

import (
	"github.com/openpeng/fool-pay/notify"
)

type ReqData struct {
	Body           string `json:"body" form:"body"  xml:"body"`
	Subject        string `json:"subject" form:"subject"  xml:"subject"`
	Openid         string `json:"openid" form:"openid"  xml:"openid"`
	OrderNo        string `json:"order_no" form:"order_no"  xml:"order_no"`
	TimeoutExpress int64  `json:"timeout_express" form:"timeout_express"  xml:"timeout_express"`
	Amount         string `json:"amount" form:"amount"  xml:"amount"`
	ReturnParam    string `json:"return_param" form:"return_param"  xml:"return_param"`
	// 支付宝公有
	GoodsType int8   `json:"goods_type" form:"goods_type"  xml:"goods_type"`
	StoreId   string `json:"store_id" form:"store_id"  xml:"store_id"`

	// 条码支付
	OperatorId    string `json:"operator_id" form:"operator_id"  xml:"operator_id"`
	TerminalId    string `json:"terminal_id" form:"terminal_id"  xml:"terminal_id"`
	AlipayStoreId string `json:"alipay_store_id" form:"alipay_store_id"  xml:"alipay_store_id"`
	Scene         string `json:"scene" form:"scene"  xml:"scene"`
	AuthCode      string `json:"auth_code" form:"auth_code"  xml:"auth_code"`

	// Web支付
	QrMod     string `json:"qr_mod" form:"qr_mod"  xml:"qr_mod"`
	Paymethod string `json:"paymethod" form:"paymethod"  xml:"paymethod"`

	ClientIp  string `json:"client_ip" form:"client_ip"  xml:"client_ip"`
	ProductId string `json:"product_id" form:"product_id"  xml:"product_id"`
}

type ResData struct {
	notify.NotifyProcessData
}

type BaseData struct {
	ReqData ReqData
	ResData ResData
}
