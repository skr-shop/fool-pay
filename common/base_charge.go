package common

// 支付主流程：
// 1. 初始化数据，设置各种支付要用到的属性
// 2. 检测一些配置的数据，是否有问题
// 3. 构建要请求三方的数据【不管是否要请求】
// 3.1 构建初始请求数据
// 3.2 构建签名
// 3.3 拼接完成
// 4. 发起请求，获取返回的数据
// 5. 利用返回的数据，构建返回给前端的数据

type ChargeClientInterface interface {
	Charge(data ReqData) interface{}
}

type ChargeClient struct {
	ConfigData BaseConfig
	ReqData    ReqData
	ChargeClientInterface
}

func NewChargeClient(configData BaseConfig, intface ChargeClientInterface) *ChargeClient {
	return &ChargeClient{
		ConfigData:            configData,
		ChargeClientInterface: intface,
	}
}
