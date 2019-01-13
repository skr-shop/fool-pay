package wx

import (
	"strconv"
	"time"

	"github.com/openpeng/fool-pay/common"
	wxbase "github.com/openpeng/fool-pay/common/wx"
	"github.com/openpeng/fool-pay/common/wx/data"
)

type ClientInterface interface {
	BuildData() string
}

type ChargeClient struct {
	*wxbase.ChargeClient
	ClientInterface
}

func NewChargeClient(configData common.BaseConfig, intface interface{}) *ChargeClient {
	var tmp = &ChargeClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.ChargeClient = wxbase.NewChargeClient(configData, intface)
	return tmp
}

func (pc *ChargeClient) GetSignType() string {
	return "MD5"
}

func (pc *ChargeClient) BuildResData() interface{} {
	var resPar = data.ResCharge{
		AppID:     pc.WeChatReResult.AppID,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  pc.WeChatReResult.NonceStr,
		Package:   "prepay_id=" + pc.WeChatReResult.PrepayID,
		SignType:  pc.GetSignType(),
		Sign:      "",
	}
	var allParams = map[string]string{
		"appId":     resPar.AppID,
		"timeStamp": resPar.TimeStamp,
		"nonceStr":  resPar.NonceStr,
		"package":   resPar.Package,
		"signType":  resPar.SignType,
	}
	resPar.Sign, _ = pc.GetSign(allParams)
	allParams["paySign"] = resPar.Sign
	return allParams
}
