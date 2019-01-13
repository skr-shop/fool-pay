package charge

import (
	aliCharge "github.com/openpeng/fool-pay/client/charge/ali"
	wxCharge "github.com/openpeng/fool-pay/client/charge/wx"
	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/constant"
	"github.com/openpeng/fool-pay/errors"
)

var supportChannel = []constant.PayChannel{
	constant.WX_CHANNEL_PUB,
	constant.WX_CHANNEL_APP,
	constant.WX_CHANNEL_LITE,
	constant.WX_CHANNEL_QR,
	constant.ALI_CHANNEL_WAP,
	constant.ALI_CHANNEL_APP,
	constant.ALI_CHANNEL_OLD_QUICK_WAP,
	constant.ALI_CHANNEL_OLD_QUICK_WEB,
}

func Run(channel constant.PayChannel, config common.BaseConfig, data common.ReqData) (retdata interface{}, iswrong errors.PayError) {
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
	retdata = handle.Charge(data)
	return
}

//数据绑定
func getHandle(channel constant.PayChannel, config common.BaseConfig) common.ChargeClientInterface {
	var handle interface{}
	switch channel {
	case constant.WX_CHANNEL_PUB:
		fallthrough
	case constant.WX_CHANNEL_LITE:
		handle = wxCharge.NewWxPubCharge(config)
		break
	case constant.WX_CHANNEL_APP:
		handle = wxCharge.NewWxAppCharge(config)
		break
	case constant.ALI_CHANNEL_WAP:
		handle = aliCharge.NewAliWapCharge(config)
		break
	case constant.ALI_CHANNEL_APP:
		handle = aliCharge.NewAliAppCharge(config)
		break
	case constant.ALI_CHANNEL_OLD_QUICK_WEB:
		handle = aliCharge.NewAliOldWebCharge(config)
		break
	case constant.ALI_CHANNEL_OLD_QUICK_WAP:
		handle = aliCharge.NewAliOldWapCharge(config)
		break
	default:
		errors.ThrewError(errors.NO_SUPPORT_CHANNEL)
	}
	return handle.(common.ChargeClientInterface)
}
