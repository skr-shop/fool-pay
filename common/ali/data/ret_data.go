package data

// AliResResult 微信支付返回
type AliResResult struct {
	PayUrl string `json:"pay_url" form:"pay_url"  xml:"pay_url"`
}

// AliPayResult 支付宝支付结果回调 https://docs.open.alipay.com/194/103296/
type AliPayResult struct {
	NotifyTime       string  `json:"notify_time"`
	NotifyType       string  `json:"notify_type"`
	NotifyID         string  `json:"notify_id"`
	SignType         string  `json:"sign_type"`
	Sign             string  `json:"sign"`
	OutTradeNum      string  `json:"out_trade_no"`
	Subject          string  `json:"subject"`
	PayMentType      string  `json:"payment_type"`
	TradeNo          string  `json:"trade_no"`
	TradeStatus      string  `json:"trade_status"`
	GmtPayMent       string  `json:"gmt_payment"`
	GmtClose         string  `json:"gmt_close"`
	SellerEmail      string  `json:"seller_email"`
	BuyerEmail       string  `json:"buyer_email"`
	SellerID         string  `json:"seller_id"`
	BuyerID          string  `json:"buyer_id"`
	Price            string  `json:"price"`
	TotalFee         float64 `json:"total_fee"`
	Quantity         string  `json:"quantity"`
	Body             string  `json:"body"`
	Discount         string  `json:"discount"`
	IsTotalFeeAdjust string  `json:"is_total_fee_adjust"`
	UseCoupon        string  `json:"use_coupon"`
	RefundStatus     string  `json:"refund_status"`
	GmtRefund        string  `json:"gmt_refund"`
	GmtCreate        string  `json:"gmt_create"`
	PassbackParams   string  `json:"passback_params"`
}

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
