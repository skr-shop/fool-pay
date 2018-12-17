package data

// AliResResult 微信支付返回
type AliResResult struct {
	PayUrl string `json:"pay_url" form:"pay_url"  xml:"pay_url"`
}

// // WechatBaseResult 基本信息
// type WechatBaseResult struct {
// 	ReturnCode string `json:"return_code" form:"return_code"  xml:"return_code"`
// 	ReturnMsg  string `json:"return_msg" form:"return_msg"  xml:"return_msg"`
// }

// // WechatReturnData 返回通用数据
// type WechatReturnData struct {
// 	AppID      string `json:"appid" form:"appid"  xml:"appid,omitempty"`
// 	MchID      string `json:"mch_id" form:"mch_id"  xml:"mch_id,omitempty"`
// 	DeviceInfo string `json:"device_info" form:"device_info"  xml:"device_info,omitempty"`
// 	NonceStr   string `json:"nonce_str" form:"nonce_str"  xml:"nonce_str,omitempty"`
// 	Sign       string `json:"sign" form:"sign"  xml:"sign,omitempty"`
// 	ResultCode string `json:"result_code" form:"result_code"  xml:"result_code,omitempty"`
// 	ErrCode    string `json:"err_code" form:"err_code"  xml:"err_code,omitempty"`
// 	ErrCodeDes string `json:"err_code_des" form:"err_code_des"  xml:"err_code_des,omitempty"`
// }

// // WechatResultData 结果通用数据
// type WechatResultData struct {
// 	OpenID        string `json:"openid" form:"openid"  xml:"openid,omitempty"`
// 	IsSubscribe   string `json:"is_subscribe" form:"is_subscribe"  xml:"is_subscribe,omitempty"`
// 	TradeType     string `json:"trade_type" form:"trade_type"  xml:"trade_type,omitempty"`
// 	BankType      string `json:"bank_type" form:"bank_type"  xml:"bank_type,omitempty"`
// 	FeeType       string `json:"fee_type" form:"fee_type"  xml:"fee_type,omitempty"`
// 	TotalFee      int64  `json:"total_fee" form:"total_fee"  xml:"total_fee,omitempty"`
// 	CashFeeType   string `json:"cash_fee_type" form:"cash_fee_type"  xml:"cash_fee_type,omitempty"`
// 	CashFee       int64  `json:"cash_fee" form:"cash_fee"  xml:"cash_fee,omitempty"`
// 	TransactionID string `json:"transaction_id" form:"transaction_id"  xml:"transaction_id,omitempty"`
// 	OutTradeNO    string `json:"out_trade_no" form:"out_trade_no"  xml:"out_trade_no,omitempty"`
// 	Attach        string `json:"attach" form:"attach"  xml:"attach,omitempty"`
// 	TimeEnd       string `json:"time_end" form:"time_end"  xml:"time_end,omitempty"`
// }

// type WeChatPayResult struct {
// 	WechatBaseResult
// 	WechatReturnData
// 	WechatResultData
// }

// type WeChatQueryResult struct {
// 	WechatBaseResult
// 	WechatReturnData
// 	WechatResultData
// 	TradeState     string `json:"trade_state" form:"trade_state"  xml:"trade_state"`
// 	TradeStateDesc string `json:"trade_state_desc" form:"trade_state_desc"  xml:"trade_state_desc"`
// }

type ResCharge struct {
	AppID     string
	TimeStamp string
	NonceStr  string
	Package   string
	SignType  string
	Sign      string
}

type ResAppCharge struct {
	AppID     string `json:"appid"`
	Partnerid string `json:"partnerid"`
	Prepayid  string `json:"prepayid"`
	Package   string `json:"package"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	SignType  string `json:"sign_type"`
	Sign      string `json:"sign"`
}
