package alinotify

import (
	"strings"

	"github.com/openpeng/fool-pay/common"
	alibase "github.com/openpeng/fool-pay/common/ali"
	"github.com/openpeng/fool-pay/errors"
)

type ClientInterface interface {
	BuildResData() string
}

type NotifyClient struct {
	*alibase.NotifyClient
	ClientInterface
}

func NewNotifyClient(configData common.BaseConfig, intface interface{}) *NotifyClient {
	var tmp = &NotifyClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.NotifyClient = alibase.NewNotifyClient(configData, intface.(alibase.NotifyClientInterface))
	return tmp
}

func (pc *NotifyClient) GetSignType() string {
	switch strings.ToUpper(pc.ConfigData.ConfigAliData.SignType) {
	case "MD5":
		return "MD5"
	case "RSA2":
		return "RSA2"
	case "RSA":
		return "RSA"
	default:
		errors.ThrewError(errors.NO_SUPPORT_SIGNTYPE)
	}
	return "RSA"
}
