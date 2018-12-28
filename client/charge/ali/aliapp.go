package ali

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/errors"

	"github.com/openpeng/fool-pay/constant"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/ali/data"
)

type AliAppCharge struct {
	*ChargeClient
}

func NewAliAppCharge(configData common.BaseConfig) *AliAppCharge {
	temp := &AliAppCharge{}
	temp.ChargeClient = NewChargeClient(configData, temp)
	return temp
}

func (wpc *AliAppCharge) BuildData() string {
	wccc := wpc.ChargeClient.ChargeClient.ConfigData.ConfigAliData
	var cpd = data.ChargePub{
		AppId:      wccc.AppId,
		Method:     string(constant.APP_PAY_METHOD),
		Format:     "JSON",
		ReturnUrl:  wccc.ReturnUrl,
		Charset:    "UTF-8",
		SignType:   wpc.GetSignType(),
		Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
		Version:    "1.0",
		NotifyUrl:  wccc.NotifyUrl,
		BizContent: wpc.GetBizContent(),
	}

	b, _ := json.Marshal(cpd)
	var allParams = make(map[string]string, 0)
	json.Unmarshal(b, &allParams)
	sign, _ := wpc.ChargeClient.GetSign(allParams)
	allParams["sign"] = sign
	wpc.AliResResult.PayUrl = wpc.ToURL(allParams)
	return wpc.AliResResult.PayUrl
}

// ToURL
func (wpc *AliAppCharge) ToURL(m map[string]string) string {
	var buf []string
	for k, v := range m {
		if v == "" {
			continue
		}
		buf = append(buf, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	return fmt.Sprintf("%s", strings.Join(buf, "&"))
}

func (wpc *AliAppCharge) BuildResData() interface{} {
	return wpc.AliResResult
}

func (wpc *AliAppCharge) GetBizContent() string {
	wccc := wpc.ChargeClient.ChargeClient.ConfigData.ConfigAliData
	wcr := wpc.ChargeClient.ReqData
	d := data.BizContent{
		Body:               wcr.Body,
		Subject:            wcr.Subject,
		OutTradeNo:         wcr.OrderNo,
		TotalAmount:        fmt.Sprintf("%.2f", wcr.Amount),
		SellerId:           wccc.Partner,
		ProductCode:        "QUICK_WAP_PAY",
		GoodsType:          wcr.GoodsType,
		PassbackParams:     wcr.ReturnParam,
		DisablePayChannels: strings.Join(wccc.LimitPay, ","),
		StoreId:            wcr.StoreId,
	}
	b, err := json.Marshal(d)
	if err != nil {
		errors.ThrewError(errors.PAY_DATA_TRS_ERROR)
		//出错了
	}
	return string(b)
}

func (wpc *AliAppCharge) GetSignType() string {
	switch strings.ToUpper(wpc.ConfigData.ConfigAliData.SignType) {
	case "RSA2":
		return "RSA2"
	case "RSA":
		return "RSA"
	}
	return "RSA"
}
