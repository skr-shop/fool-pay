package ali

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/errors"
)

type AliOldWebCharge struct {
	*ChargeClient
}

func NewAliOldWebCharge(configData common.BaseConfig) *AliOldWebCharge {
	temp := &AliOldWebCharge{}
	temp.ChargeClient = NewChargeClient(configData, temp)
	return temp
}

//参考文档：https://docs.open.alipay.com/62/104743/

func (wpc *AliOldWebCharge) BuildData() string {
	wccc := wpc.ChargeClient.ChargeClient.ConfigData.ConfigAliData
	wcr := wpc.ChargeClient.ReqData

	var allParams = map[string]string{
		"service":           "create_direct_pay_by_user",
		"partner":           wccc.Partner,
		"_input_charset":    "utf-8",
		"sign_type":         wpc.GetSignType(),
		"sign":              "",
		"notify_url":        wccc.NotifyUrl,
		"return_url":        wccc.ReturnUrl,
		"out_trade_no":      wcr.OrderNo,
		"subject":           wcr.Subject,
		"payment_type":      "1",
		"total_fee":         fmt.Sprintf("%.2f", wcr.Amount),
		"seller_id":         wccc.Partner,
		"seller_email":      wccc.SellerEmail,
		"body":              wcr.Body,
		"disable_paymethod": strings.Join(wccc.LimitPay, "^"),
		"exter_invoke_ip":   wcr.ClientIp,
		"it_b_pay":          fmt.Sprintf("%dm", int((wcr.TimeoutExpress-time.Now().Unix())/60)),
		"goods_type":        fmt.Sprintf("%d", wcr.GoodsType),
		"extend_param":      wcr.ReturnParam,
	}
	allParams["sign"], _ = wpc.ChargeClient.GetSign(allParams)
	wpc.AliResResult.PayUrl = wpc.ToURL(allParams)
	return wpc.AliResResult.PayUrl
}

// ToURL
func (wpc *AliOldWebCharge) ToURL(m map[string]string) string {
	var buf []string
	for k, v := range m {
		if v == "" {
			continue
		}
		buf = append(buf, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	return fmt.Sprintf("http://mapi.alipay.com/gateway.do?%s", strings.Join(buf, "&"))
}

func (wpc *AliOldWebCharge) BuildResData() interface{} {
	return wpc.AliResResult
}

func (wpc *AliOldWebCharge) GetSignType() string {
	switch strings.ToUpper(wpc.ConfigData.ConfigAliData.SignType) {
	case "MD5":
		return "MD5"
	case "RSA":
		return "RSA"
	}
	errors.ThrewError(errors.PAY_WAY_NO_SIGN_TYPE)
	return ""
}

func (wpc *AliOldWebCharge) IsOldPay() bool {
	return true
}
