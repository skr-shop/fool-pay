package charge

import (
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
}

func Run(channel constant.PayChannel, config common.BaseConfig, data common.ReqData) (retdata interface{}, iswrong errors.PayError) {
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
	handle.Charge(data)
	retdata = handle.Send()
	return
}

//数据绑定
func getHandle(channel constant.PayChannel, config common.BaseConfig) *common.ChargeClient {
	var handle common.ChargeClient
	switch channel {
	// case constant.WX_CHANNEL_APP:
	// 	ser := wxCharge.NewWxAppCharge(config)
	// 	handle = common.NewChargeClient(ser)
	// 	break
	// 小程序支付与公众号支付一样，仅仅是客户端的调用方式不同
	case constant.WX_CHANNEL_PUB:
		fallthrough
	case constant.WX_CHANNEL_LITE:
		ser := wxCharge.NewWxPubCharge(config)
		handle = *common.NewChargeClient(ser)
		break
	// case Config::WX_CHANNEL_WAP:
	//     $this->channel = new WxWapCharge($config);
	//     break;
	// case Config::WX_CHANNEL_QR:
	//     $this->channel = new WxQrCharge($config);
	//     break;
	// case Config::WX_CHANNEL_BAR:
	//     $this->channel = new WxBarCharge($config);
	//     break;
	// case Config::ALI_CHANNEL_WAP:
	//     $this->channel = new AliWapCharge($config);
	//     break;
	// case Config::ALI_CHANNEL_APP:
	//     $this->channel = new AliAppCharge($config);
	//     break;
	// case Config::ALI_CHANNEL_WEB:
	//     $this->channel = new AliWebCharge($config);
	//     break;
	// case Config::ALI_CHANNEL_QR:
	//     $this->channel = new AliQrCharge($config);
	//     break;
	// case Config::ALI_CHANNEL_BAR:
	//     $this->channel = new AliBarCharge($config);
	//     break;
	default:
		errors.ThrewError(errors.NO_SUPPORT_CHANNEL)
	}
	return &handle
}
