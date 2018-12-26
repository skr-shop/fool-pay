package notify

// 处理成功看这里 json目前是支付宝对应的，xml目前是微信
type NotifyProcessData struct {
	TradeStatus string      `json:"trade_status,omitempty" xml:"result_code,omitempty"` //交易状态only SUCCESS  CLOSED
	Amount      float64     `json:"total_amount,omitempty" xml:"total_fee,omitempty"`   //总金额，单位元
	Attach      string      `json:"passback_params,omitempty" xml:"attach,omitempty"`   //公共回传参数
	BuyerId     string      `json:"buyer_id" xml:"openid"`                              //买家账号
	OrderNo     string      `json:"order_no" xml:"transaction_id"`                      //三方交易凭证号
	OutTradeNo  string      `json:"out_trade_no" xml:"out_trade_no"`                    //原支付请求的商户订单号
	PayTime     int64       `json:"gmt_payment" xml:"time_end"`                         //支付时间 yyyy-MM-dd HH:mm:ss
	Origin      interface{} `json:"origin" xml:"origin"`                                //原始数据
}

type NotifyInterface interface {
	NotifyProcess(data NotifyProcessData) bool
}
