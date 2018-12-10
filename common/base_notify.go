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

func NewNotifyClient(configData BaseConfig, intface NotifyClientInterface) *NotifyClient {
	return &NotifyClient{
		ConfigData:            configData,
		NotifyClientInterface: intface,
	}
}
