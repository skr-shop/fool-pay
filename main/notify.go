package main

import (
	"fmt"

	notifyInt "github.com/openpeng/fool-pay/notify"

	"github.com/openpeng/fool-pay/client/notify"
	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/constant"
)

type Noti struct {
	notifyInt.NotifyInterface
}

func (nt Noti) NotifyProcess(ntp notifyInt.NotifyProcessData) bool {
	fmt.Println(ntp)
	return true
}

func main() {
	var config = common.BaseConfig{
		ConfigWxData: common.ConfigWxData{
			AppId:       "",
			MchId:       "",
			Md5Key:      "",
			AppCertPem:  "app_cert_pem",
			AppKeyPem:   "app_key_pem",
			SignType:    "MD5",
			LimitPay:    []string{},
			FeeType:     "CNY",
			NotifyUrl:   "",
			RedirectUrl: "redirect_url",
		},
		ConfigPubData: common.ConfigPubData{
			ReturnRaw: false,
		},
	}
	var st = `
	<xml>
	<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
	<attach><![CDATA[支付测试]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[10000100]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[1409811653]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
	<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>1</total_fee>
  <coupon_fee_0><![CDATA[10]]></coupon_fee_0>
  <coupon_count><![CDATA[1]]></coupon_count>
  <coupon_type><![CDATA[CASH]]></coupon_type>
  <coupon_id><![CDATA[10000]]></coupon_id> 
	<trade_type><![CDATA[JSAPI]]></trade_type>
	<transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
  </xml>
  `
	dd := []byte(st)
	np := Noti{}
	i, _ := notify.Run(constant.WX_CHARGE, config, dd, np)
	// WeChatResult 微信支付返回
	// n := i.(data.ResCharge)
	// s := fmt.Sprintf("https://t-mtiku.gaodun.com/activity/pay/1?appId=%s&timeStamp=%s&nonceStr=%s&package=%s&signType=%s&paySign=%s", n.AppID, n.TimeStamp, n.NonceStr, n.Package, n.SignType, n.Sign)
	//  appId  timeStamp    nonceStr   package   signType   paySign

	// i = charge.Run(constant.WX_CHANNEL_APP, config, reqdata)
	fmt.Println(i)
}
