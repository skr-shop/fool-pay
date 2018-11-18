package pay

// 支付接口
type Pay interface {
	send(params []string)
}
