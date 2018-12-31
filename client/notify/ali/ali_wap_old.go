package alinotify

import (
	"github.com/openpeng/fool-pay/common"
)

type AliOldNotify struct {
	*NotifyClient
}

func NewAliOldNotify(config common.BaseConfig) *AliOldNotify {
	temp := &AliOldNotify{}
	temp.NotifyClient = NewNotifyClient(config, temp)
	return temp
}

func (wn *AliOldNotify) BuildResData() string {
	return "success"
}

func (wn *AliOldNotify) BuildData() string {
	return "success"
}
func (wn *AliOldNotify) IsOldPay() bool {
	return true
}
