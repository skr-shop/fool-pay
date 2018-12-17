package ali

import (
	"github.com/openpeng/fool-pay/common"
	alibase "github.com/openpeng/fool-pay/common/ali"
)

type ClientInterface interface {
	BuildData() string
}

type ChargeClient struct {
	*alibase.ChargeClient
	ClientInterface
}

func NewChargeClient(configData common.BaseConfig, intface interface{}) *ChargeClient {
	var tmp = &ChargeClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.ChargeClient = alibase.NewChargeClient(configData, intface)
	return tmp
}

func (pc *ChargeClient) GetSignType() string {
	return "RSA"
}
