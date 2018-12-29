package alinotify

import (
	"github.com/openpeng/fool-pay/common"
)

type AliNotify struct {
	*NotifyClient
}

func NewAliNotify(config common.BaseConfig) *AliNotify {
	temp := &AliNotify{}
	temp.NotifyClient = NewNotifyClient(config, temp)
	return temp
}

func (wn *AliNotify) BuildResData() string {
	return "success"
}

func (wpc *AliNotify) BuildData() string {
	return "success"
}
