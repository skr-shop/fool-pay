package data

type ChargePub struct {
	AppId      string `json:"app_id,omitempty" form:"app_id"  xml:"app_id,omitempty,cdata"`
	Method     string `json:"method,omitempty" form:"method"  xml:"method,omitempty,cdata"`
	Format     string `json:"format,omitempty" form:"format"  xml:"format,omitempty,cdata"`
	ReturnUrl  string `json:"return_url,omitempty" form:"return_url"  xml:"return_url,omitempty,cdata"`
	Charset    string `json:"charset,omitempty" form:"charset"  xml:"charset,omitempty,cdata"`
	SignType   string `json:"sign_type,omitempty" form:"sign_type"  xml:"sign_type,omitempty,cdata"`
	Timestamp  string `json:"timestamp,omitempty" form:"timestamp"  xml:"timestamp,omitempty,cdata"`
	Version    string `json:"version,omitempty" form:"version"  xml:"version,omitempty,cdata"`
	NotifyUrl  string `json:"notify_url,omitempty" form:"notify_url"  xml:"notify_url,omitempty,cdata"`
	BizContent string `json:"biz_content,omitempty" form:"biz_content"  xml:"biz_content,omitempty,cdata"`
}

type BizContent struct {
	Body        string `json:"body,omitempty" form:"body"  xml:"body,omitempty,cdata"`
	Subject     string `json:"subject,omitempty" form:"subject"  xml:"subject,omitempty,cdata"`
	OutTradeNo  string `json:"out_trade_no,omitempty" form:"out_trade_no"  xml:"out_trade_no,omitempty,cdata"`
	TotalAmount string `json:"total_amount,omitempty" form:"total_amount"  xml:"total_amount,omitempty,cdata"`
	SellerId    string `json:"seller_id,omitempty" form:"seller_id"  xml:"seller_id,omitempty,cdata"`
	// 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_WAP_PAY
	ProductCode        string `json:"product_code,omitempty" form:"product_code"  xml:"product_code,omitempty,cdata"`
	GoodsType          int8   `json:"goods_type,omitempty" form:"goods_type"  xml:"goods_type,omitempty,cdata"`
	PassbackParams     string `json:"passback_params,omitempty" form:"passback_params"  xml:"passback_params,omitempty,cdata"`
	DisablePayChannels string `json:"disable_pay_channels,omitempty" form:"disable_pay_channels"  xml:"disable_pay_channels,omitempty,cdata"`
	StoreId            string `json:"store_id,omitempty" form:"store_id"  xml:"store_id,omitempty,cdata"`
}
