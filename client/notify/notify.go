package notify

import (
	wxnotify "github.com/openpeng/fool-pay/client/notify/wx"
	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/constant"
	"github.com/openpeng/fool-pay/errors"
	"github.com/openpeng/fool-pay/notify"
)

var supportChannel = []constant.PayOperation{
	constant.WX_CHARGE,
	// constant.ALI_CHARGE
}

func Run(channel constant.PayOperation, config common.BaseConfig, data []byte, userNotify notify.NotifyInterface) (retdata interface{}, iswrong errors.PayError) {
	iswrong = errors.PayError{}
	defer errors.Catch(&iswrong)
	support := false
	for _, supportChannel := range supportChannel {
		if channel == supportChannel {
			support = true
		}
	}
	if !support {
		errors.ThrewError(errors.NO_SUPPORT_CHANNEL)
	}
	handle := getHandle(channel, config)
	retdata = handle.Notify(data, userNotify)
	return
}

//数据绑定
func getHandle(channel constant.PayOperation, config common.BaseConfig) common.NotifyClientInterface {
	var handle interface{}
	switch channel {
	case constant.WX_CHARGE:
		handle = wxnotify.NewWechatNotify(config)
		break
	default:
		errors.ThrewError(errors.NO_SUPPORT_CHANNEL)
	}
	return handle.(common.NotifyClientInterface)
}
