package wxnotify

import (
	"github.com/openpeng/fool-pay/common"
	wxbase "github.com/openpeng/fool-pay/common/wx"
)

type ClientInterface interface {
	BuildResData() string
}

type NotifyClient struct {
	*wxbase.NotifyClient
	ClientInterface
}

func NewNotifyClient(configData common.BaseConfig, intface interface{}) *NotifyClient {
	var tmp = &NotifyClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.NotifyClient = wxbase.NewNotifyClient(configData, intface.(wxbase.NotifyClientInterface))
	return tmp
}

func (pc *NotifyClient) GetSignType() string {
	return "MD5"
}
