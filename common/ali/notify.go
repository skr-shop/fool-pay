package ali

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/openpeng/fool-pay/constant"
	"github.com/openpeng/fool-pay/errors"
	"github.com/openpeng/fool-pay/util"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/ali/data"
	"github.com/openpeng/fool-pay/notify"
)

type NotifyClientInterface interface {
	GetSignType() string
	CheckConfig()
	IsOldPay() bool
	BuildResData() string
}

type NotifyClient struct {
	data.AliPayResult
	OriginData map[string]string
	*common.NotifyClient
	*common.SignHandle
	NotifyClientInterface
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewNotifyClient(config common.BaseConfig, intface NotifyClientInterface) *NotifyClient {
	var cc = &NotifyClient{
		NotifyClientInterface: intface,
	}
	cc.NotifyClient = common.NewNotifyClient(config, cc)
	var sc = common.SignConfig{
		SignType: constant.SignWay(strings.ToUpper(config.ConfigAliData.SignType)),
		SignKey: common.SignKey{
			PrivateKey: config.ConfigAliData.RsaPrivateKey,
			PublicKey:  config.ConfigAliData.AliPublicKey,
		},
	}
	cc.SignHandle = common.NewSignClient(sc, intface.(common.SignClientInterface))
	return cc
}

func (nc *NotifyClient) Notify(d []byte, process notify.NotifyInterface) string {
	ok := process.NotifyProcess(nc.GetNotifyData(d))
	if ok {
		return nc.NotifyClientInterface.BuildResData()
	}
	return "false"
}

func (nc *NotifyClient) GetNotifyData(b []byte) notify.NotifyProcessData {
	var notifyMapData = data.AliPayResult{}
	var signData = make(map[string]string, 0)
	json.Unmarshal(b, &notifyMapData)
	nc.AliPayResult = notifyMapData
	json.Unmarshal(b, &signData)
	nc.OriginData = signData
	if e := nc.SignHandle.CheckSign(notifyMapData.Sign); e != nil {
		errors.ThrewError(errors.SIGN_WRONG)
	}
	endTime, _ := time.Parse("2006-01-02 15:04:05", notifyMapData.NotifyTime)
	return notify.NotifyProcessData{
		Amount:      float64(notifyMapData.TotalFee / 100),
		Attach:      notifyMapData.PassbackParams,
		OrderNo:     notifyMapData.TradeNo,
		PayTime:     endTime.Unix() - 8*3600, //当前时间要减8小时
		BuyerId:     notifyMapData.BuyerID,
		OutTradeNo:  notifyMapData.OutTradeNum,
		TradeStatus: notifyMapData.TradeStatus,
		Origin:      notifyMapData,
	}
}

//检测配置
func (pc *NotifyClient) CheckConfig() {
	if pc.NotifyClientInterface.GetSignType() != "MD5" {
		pc.PrivateKey = util.Bytes2RSAPrivateKey([]byte(pc.ConfigData.ConfigAliData.RsaPrivateKey))
		pc.PublicKey = util.Bytes2RSAPublicKey([]byte(pc.ConfigData.ConfigAliData.AliPublicKey))
		if pc.ConfigData.ConfigAliData.RsaPrivateKey == "" {
			errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
		}
	} else {
		if pc.ConfigData.ConfigAliData.Key == "" {
			errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
		}
	}

}

func (nc *NotifyClient) BuildSignData() string {
	m := nc.OriginData
	delete(m, "sign")
	if nc.NotifyClientInterface.IsOldPay() {
		delete(m, "sign_type")
	}
	var data []string
	for k, v := range m {
		if v == "" {
			continue
		}
		data = append(data, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(data)
	nts := strings.Join(data, "&")

	if nc.NotifyClientInterface.GetSignType() == "MD5" {
		fmt.Println(nc.ConfigData.ConfigAliData.Key)
		return nts + nc.ConfigData.ConfigAliData.Key
	}
	return nts
}

func (nc *NotifyClient) Send() interface{} {
	return nc.NotifyClientInterface.BuildResData()
}

func (nc *NotifyClient) IsOldPay() bool {
	return false
}
