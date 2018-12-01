package main

import (
	"fmt"

	"github.com/openpeng/fool-pay/client/charge"
	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/constant"
)

func main() {
	var config = common.BaseConfig{
		ConfigWxData: common.ConfigWxData{
			AppId:       "",
			MchId:       "",
			Md5Key:      "",
			AppCertPem:  "app_cert_pem",
			AppKeyPem:   "app_key_pem",
			SignType:    "MD5",
			LimitPay:    []string{"limit_pay"},
			FeeType:     "CNY",
			NotifyUrl:   "/api/pay/xcxCallBack",
			RedirectUrl: "redirect_url",
		},
		ConfigPubData: common.ConfigPubData{
			ReturnRaw: false,
		},
	}
	reqdata := common.ReqData{}
	i := charge.Run(constant.WX_CHANNEL_PUB, config, reqdata)
	fmt.Println(i)
}
