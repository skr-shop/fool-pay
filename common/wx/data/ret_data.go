package data

// WeChatResult 微信支付返回
type WeChatReResult struct {
	ReturnCode string `json:"return_code" form:"return_code"  xml:"return_code"`
	ReturnMsg  string `json:"return_msg" form:"return_msg"  xml:"return_msg"`

	AppID      string `json:"appid" form:"appid"  xml:"appid"`
	MchID      string `json:"mch_id" form:"mch_id"  xml:"mch_id"`
	DeviceInfo string `json:"device_info" form:"device_info"  xml:"device_info"`
	NonceStr   string `json:"nonce_str" form:"nonce_str"  xml:"nonce_str"`
	Sign       string `json:"sign" form:"sign"  xml:"sign"`
	ResultCode string `json:"result_code" form:"result_code"  xml:"result_code"`
	ErrCode    string `json:"err_code" form:"err_code"  xml:"err_code"`
	ErrCodeDes string `json:"err_code_des" form:"err_code_des"  xml:"err_code_des"`

	TradeType string `json:"trade_type" form:"trade_type"  xml:"trade_type"`
	PrepayID  string `json:"prepay_id" form:"prepay_id"  xml:"prepay_id"`
	CodeURL   string `json:"code_url" form:"code_url"  xml:"code_url"`
}

// WechatBaseResult 基本信息
type WechatBaseResult struct {
	ReturnCode string `json:"return_code" form:"return_code"  xml:"return_code"`
	ReturnMsg  string `json:"return_msg" form:"return_msg"  xml:"return_msg"`
}

// WechatReturnData 返回通用数据
type WechatReturnData struct {
	AppID      string `json:"appid" form:"appid"  xml:"appid,emitempty"`
	MchID      string `json:"mch_id" form:"mch_id"  xml:"mch_id,emitempty"`
	DeviceInfo string `json:"device_info" form:"device_info"  xml:"device_info,emitempty"`
	NonceStr   string `json:"nonce_str" form:"nonce_str"  xml:"nonce_str,emitempty"`
	Sign       string `json:"sign" form:"sign"  xml:"sign,emitempty"`
	ResultCode string `json:"result_code" form:"result_code"  xml:"result_code,emitempty"`
	ErrCode    string `json:"err_code" form:"err_code"  xml:"err_code,emitempty"`
	ErrCodeDes string `json:"err_code_des" form:"err_code_des"  xml:"err_code_des,emitempty"`
}

// WechatResultData 结果通用数据
type WechatResultData struct {
	OpenID        string `json:"openid" form:"openid"  xml:"openid,emitempty"`
	IsSubscribe   string `json:"is_subscribe" form:"is_subscribe"  xml:"is_subscribe,emitempty"`
	TradeType     string `json:"trade_type" form:"trade_type"  xml:"trade_type,emitempty"`
	BankType      string `json:"bank_type" form:"bank_type"  xml:"bank_type,emitempty"`
	FeeType       string `json:"fee_type" form:"fee_type"  xml:"fee_type,emitempty"`
	TotalFee      int64  `json:"total_fee" form:"total_fee"  xml:"total_fee,emitempty"`
	CashFeeType   string `json:"cash_fee_type" form:"cash_fee_type"  xml:"cash_fee_type,emitempty"`
	CashFee       int64  `json:"cash_fee" form:"cash_fee"  xml:"cash_fee,emitempty"`
	TransactionID string `json:"transaction_id" form:"transaction_id"  xml:"transaction_id,emitempty"`
	OutTradeNO    string `json:"out_trade_no" form:"out_trade_no"  xml:"out_trade_no,emitempty"`
	Attach        string `json:"attach" form:"attach"  xml:"attach,emitempty"`
	TimeEnd       string `json:"time_end" form:"time_end"  xml:"time_end,emitempty"`
}

type WeChatPayResult struct {
	WechatBaseResult
	WechatReturnData
	WechatResultData
}

type WeChatQueryResult struct {
	WechatBaseResult
	WechatReturnData
	WechatResultData
	TradeState     string `json:"trade_state" form:"trade_state"  xml:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc" form:"trade_state_desc"  xml:"trade_state_desc"`
}

type ResCharge struct {
	AppID     string
	TimeStamp string
	NonceStr  string
	Package   string
	SignType  string
	Sign      string
}
