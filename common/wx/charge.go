package wx

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/container"

	"github.com/openpeng/fool-pay/errors"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/wx/data"
)

type ClientInterface interface {
	BuildData() string
	GetSignType() string
	CheckConfig()
	BuildResData() interface{}
}

// 调用方法的时候要指定是ClientInterface的方法，
// 这样才能实现调用到子集的数据
// 如果自身也实现了就可以实现不会报错的问题（子集继承上一层会继承到方法，实现会自动关联）
type ChargeClient struct {
	*common.ChargeClient
	*common.HttpClient
	ClientInterface
	*data.WeChatReResult
}

func NewChargeClient(configData common.BaseConfig, intface interface{}) *ChargeClient {
	var cc = &ChargeClient{
		ClientInterface: intface.(ClientInterface),
		HttpClient:      container.HttpClient,
	}
	//将继承的实现掉 配置信息和请求都是在公共的地方，可以考虑先统一实现再初始化配置，考虑先处理配置是可以判断配置有没有问题
	cc.ChargeClient = common.NewChargeClient(configData, intface.(common.ChargeClientInterface))
	return cc
}

//对外的入口，接入请求数据
func (pc *ChargeClient) Charge(data common.ReqData) interface{} {
	pc.ChargeClient.ReqData = data
	pc.ClientInterface.CheckConfig() //可以调用到子集的方法
	return pc.Send()
}

//检测配置
func (pc *ChargeClient) CheckConfig() {
	if pc.ConfigData.ConfigWxData.Md5Key == "" {
		errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
	}
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
	xmlstring := pc.BuildData() //这么调用是因为本身没实现，必定走的是子集的方法,但是还是推荐加ClientInterface
	info, err := pc.HttpClient.PostBodyXml("https://api.mch.weixin.qq.com/pay/unifiedorder", xmlstring)
	err = xml.Unmarshal(info, &pc.WeChatReResult)
	if err != nil {
		errors.ThrewMessageError(err.Error())
	}
	if pc.WeChatReResult.ReturnCode != "SUCCESS" {
		errors.ThrewMessageError(pc.WeChatReResult.ReturnMsg)
	}
	return pc.ClientInterface.BuildResData()
}

func (pc *ChargeClient) BuildResData() interface{} {
	var resPar = data.ResCharge{
		AppID:     pc.WeChatReResult.AppID,
		TimeStamp: strconv.FormatInt(time.Now().Unix(), 10),
		NonceStr:  pc.WeChatReResult.NonceStr,
		Package:   "prepay_id=" + pc.WeChatReResult.PrepayID,
		SignType:  "MD5",
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
	return resPar
}
