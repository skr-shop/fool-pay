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
	// 异常捕获的两种写法
	// 第一种
	// iswrong = errors.PayError{} //主要作用是分配内存
	// defer errors.Catch(&iswrong) //正常捕获
	// 第二种
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		switch err.(type) {
	// 		case errors.PayError:
	// 			pe := err.(errors.PayError)
	// 			iswrong.ErrorCode = pe.ErrorCode
	// 			iswrong.Message = pe.Message
	// 		}
	// 	}
	// }()
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
func getHandle(channel constant.PayOperation, config common.BaseConfig) *common.NotifyClient {
	var handle common.NotifyClient
	switch channel {
	case constant.WX_CHARGE:
		ser := wxnotify.NewWechatNotify(config)
		handle = *common.NewNotifyClient(ser)
		break
	default:
		errors.ThrewError(errors.NO_SUPPORT_CHANNEL)
	}
	return &handle
}
