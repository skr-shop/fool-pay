package common

type ChargeClientInterface interface {
	Send() interface{}
	BuildData() interface{}
	Charge(data ReqData) interface{}
	GetSign(m map[string]string) (string, error)
}

type ChargeClient struct {
	ConfigData BaseConfig
	ReqData    ReqData
	ChargeClientInterface
}

func NewChargeClient(intface ChargeClientInterface) *ChargeClient {
	return &ChargeClient{
		ChargeClientInterface: intface,
	}
}
