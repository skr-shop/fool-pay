package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/wx"
	"github.com/openpeng/fool-pay/common/wx/data"
)

type WxPubCharge struct {
	ChargeClient *common.ChargeClient
}

func NewWxPubCharge(config common.BaseConfig) *common.ChargeClient {
	temp := &WxPubCharge{}
	wxc := wx.NewChargeClient(temp)
	wxc.ConfigData = config
	temp.ChargeClient = wxc
	return wxc
}

func (wpc *WxPubCharge) BuildData() string {
	wccc := wpc.ChargeClient.ConfigData.ConfigWxData
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
		TradeType:      "JSAPI",
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
