package wx

import (
	"github.com/openpeng/fool-pay/common"
	wxbase "github.com/openpeng/fool-pay/common/wx"
)

type ClientInterface interface {
	BuildData() string
}

type ChargeClient struct {
	*wxbase.ChargeClient
	ClientInterface
}

func NewChargeClient(configData common.BaseConfig, intface interface{}) *ChargeClient {
	var tmp = &ChargeClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.ChargeClient = wxbase.NewChargeClient(configData, intface)
	return tmp
}

func (pc *ChargeClient) GetSignType() string {
	return "MD5"
}
