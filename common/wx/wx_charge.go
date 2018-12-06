package wx

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	"github.com/openpeng/fool-pay/errors"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/wx/data"
)

type ClientInterface interface {
	BuildData() string
}

type ChargeClient struct {
	*common.ChargeClient
	This ClientInterface
}

func NewChargeClient(intface ClientInterface) *common.ChargeClient {
	var cc = &ChargeClient{
		This: intface,
	}
	cc.ChargeClient = common.NewChargeClient(cc)
	return cc.ChargeClient
}

func (pc *ChargeClient) Charge(data common.ReqData) interface{} {
	pc.ReqData = data
	return "准备拼接信息"
}

func (pc *ChargeClient) SetSign() interface{} {
	return "签名有没有?"
}

// GetSign 产生签名
func (pc *ChargeClient) GetSign(m map[string]string) (string, error) {
	delete(m, "sign")
	var signData = make([]string, 0)
	for k, v := range m {
		if v != "" {
			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(signData)
	signStr := strings.Join(signData, "&")
	signStr = signStr + "&key=" + pc.ConfigData.ConfigWxData.Md5Key
	fmt.Println(signStr)
	c := md5.New()
	_, err := c.Write([]byte(signStr))
	if err != nil {
		return "", err
	}
	signByte := c.Sum(nil)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(fmt.Sprintf("%x", signByte)), nil
}

func (pc *ChargeClient) Send() interface{} {
	xmlstring := pc.This.BuildData()
	info, err := common.HttpClientHandle.PostBodyXml("https://api.mch.weixin.qq.com/pay/unifiedorder", xmlstring)
	var xmlRe data.WeChatReResult
	err = xml.Unmarshal(info, &xmlRe)
	if err != nil {
		fmt.Println(err)
	}
	if xmlRe.ReturnCode != "SUCCESS" {
		errors.ThrewMessageError(xmlRe.ReturnMsg)
	}
	return xmlRe
}
