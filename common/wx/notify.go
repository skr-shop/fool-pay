package wx

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"

	"github.com/openpeng/fool-pay/errors"
	"github.com/openpeng/fool-pay/util"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/wx/data"
	"github.com/openpeng/fool-pay/notify"
)

type NotifyClientInterface interface {
	BuildData() string
}

type NotifyClient struct {
	*common.NotifyClient
	NotifyClientInterface
}

func NewNotifyClient(config common.BaseConfig, intface NotifyClientInterface) *NotifyClient {
	var cc = &NotifyClient{
		NotifyClientInterface: intface,
	}
	cc.NotifyClient = common.NewNotifyClient(config, cc)
	return cc
}

func (nc *NotifyClient) Notify(d []byte, process notify.NotifyInterface) string {
	ok := process.NotifyProcess(nc.GetNotifyData(d))
	if ok {
		return "ok"
	}
	return "false"
}

func (nc *NotifyClient) GetNotifyData(b []byte) notify.NotifyProcessData {
	var notifyMapData = data.WeChatPayResult{}
	signData, _ := util.XmlToMap(b)
	xml.Unmarshal(b, &notifyMapData)
	nc.CheckSign(signData, notifyMapData.Sign)
	return notify.NotifyProcessData{
		Amount:      float32(notifyMapData.TotalFee) / 100,
		Attach:      notifyMapData.Attach,
		OrderNo:     notifyMapData.OutTradeNO,
		PayTime:     notifyMapData.TimeEnd,
		BuyerId:     notifyMapData.OpenID,
		OutTradeNo:  notifyMapData.OutTradeNO,
		TradeStatus: notifyMapData.ReturnMsg,
		Origin:      notifyMapData,
	}
}

func (nc *NotifyClient) CheckSign(od map[string]string, sign string) bool {
	ns, err := nc.GetSign(od)
	if err != nil || ns != sign {
		errors.ThrewError(errors.SIGN_WRONG)
	}
	return true
}

// GetSign 产生签名
func (nc *NotifyClient) GetSign(m map[string]string) (string, error) {
	delete(m, "sign")
	delete(m, "key")
	var signData = make([]string, 0)
	for k, v := range m {
		if v != "" {
			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(signData)
	signStr := strings.Join(signData, "&")
	signStr = signStr + "&key=" + nc.ConfigData.ConfigWxData.Md5Key
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

//检测配置
func (pc *NotifyClient) CheckConfig() {
	if pc.ConfigData.ConfigWxData.Md5Key == "" {
		errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
	}
}
