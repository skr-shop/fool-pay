package common

import "github.com/openpeng/fool-pay/notify"

type NotifyClientInterface interface {
	Notify(d []byte, process notify.NotifyInterface) bool
}

type NotifyClient struct {
	ConfigData BaseConfig
	ResData    ResData
	NotifyClientInterface
}

func NewNotifyClient(intface NotifyClientInterface) *NotifyClient {
	return &NotifyClient{
		NotifyClientInterface: intface,
	}
}
