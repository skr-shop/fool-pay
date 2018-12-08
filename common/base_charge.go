package common

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
