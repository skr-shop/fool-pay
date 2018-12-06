package main

import (
	"fmt"
	"time"

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
	reqdata := common.ReqData{
		Body:           "test",
		Subject:        "test",
		Openid:         "oH6wZt5ui06BX2WLcLc_-yyTygR4",
		OrderNo:        "xcx1111111111111111",
		TimeoutExpress: 600 + time.Now().Unix(),
		Amount:         "1",
		ReturnParam:    "963",
		GoodsType:      1,
		StoreId:        "",
		OperatorId:     "",
		TerminalId:     "",
		AlipayStoreId:  "",
		Scene:          "bar_code",
		AuthCode:       "1231212232323123123",
		QrMod:          "",
		Paymethod:      "creditPay",
		ClientIp:       "127.0.0.1",
		ProductId:      "1",
	}
	i, e := charge.Run(constant.WX_CHANNEL_PUB, config, reqdata)
	// WeChatResult 微信支付返回
	fmt.Println(i)
	fmt.Println(e)
	// i = charge.Run(constant.WX_CHANNEL_APP, config, reqdata)
	// fmt.Println(i)
}
