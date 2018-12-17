package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/wx/data"
)

type WxAppCharge struct {
	*ChargeClient
}

func NewWxAppCharge(configData common.BaseConfig) *WxAppCharge {
	temp := &WxAppCharge{}
	temp.ChargeClient = NewChargeClient(configData, temp)
	return temp
}

func (wpc *WxAppCharge) BuildData() string {
	wccc := wpc.ChargeClient.ChargeClient.ConfigData.ConfigWxData
	wcr := wpc.ChargeClient.ReqData
	var cpd = data.ChargePub{
		Appid:          wccc.AppId,
		MchId:          wccc.MchId,
		SignType:       wccc.SignType,
		FeeType:        wccc.FeeType,
		NotifyUrl:      wccc.NotifyUrl,
		LimitPay:       strings.Join(wccc.LimitPay, ","),
		DeviceInfo:     "WEB",
		Body:           wcr.Body,
		Attach:         wcr.ReturnParam,
		OutTradeNo:     wcr.OrderNo,
		TimeExpire:     time.Unix(wcr.TimeoutExpress, 0).Format("20060102150405"),
		Openid:         wcr.Openid,
		TotalFee:       wcr.Amount,
		TradeType:      "APP",
		SpbillCreateIp: wcr.ClientIp,
		TimeStart:      time.Now().Format("20060102150405"),
		NonceStr:       "11221315456",
	}

	b, _ := json.Marshal(cpd)
	var allParams = make(map[string]string, 0)
	json.Unmarshal(b, &allParams)
	sign, _ := wpc.ChargeClient.GetSign(allParams)
	// 转出xml结构
	allParams["sign"] = sign
	buf := bytes.NewBufferString("")
	for k, v := range allParams {
		buf.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", k, v, k))
	}
	xmlStr := fmt.Sprintf("<xml>%s</xml>", buf.String())
	return xmlStr
}

func (pc *ChargeClient) BuildResData() interface{} {
	var resPar = data.ResAppCharge{
		AppID:     pc.WeChatReResult.AppID,
		Partnerid: pc.ConfigData.ConfigWxData.MchId,
		Prepayid:  pc.WeChatReResult.PrepayID,
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		Noncestr:  pc.WeChatReResult.NonceStr,
		Package:   "Sign=WXPay",
	}
	var allParams = map[string]string{
		"appid":     resPar.AppID,
		"timestamp": resPar.Timestamp,
		"noncestr":  resPar.Noncestr,
		"package":   resPar.Package,
		"prepayid":  resPar.Prepayid,
		"partnerid": resPar.Partnerid,
	}
	resPar.Sign, _ = pc.GetSign(allParams)
	return resPar
}
