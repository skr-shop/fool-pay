package ali

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

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
	*common.NotifyClient
	NotifyClientInterface
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
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
	nc.CheckSign(signData, notifyMapData.Sign)
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

func (nc *NotifyClient) CheckSign(od map[string]string, sign string) bool {
	return true
	ns, err := nc.GetSign(od)
	if err != nil || ns != sign {
		errors.ThrewError(errors.SIGN_WRONG)
	}
	return true
}

// // GetSign 产生签名
// func (nc *NotifyClient) GetSign(m map[string]string) (string, error) {
// 	delete(m, "sign")
// 	delete(m, "sign_type")
// 	var signData = make([]string, 0)
// 	for k, v := range m {
// 		if v != "" {
// 			signData = append(signData, fmt.Sprintf("%s=%s", k, v))
// 		}
// 	}
// 	sort.Strings(signData)
// 	signStr := strings.Join(signData, "&")
// 	signStr = signStr + nc.ConfigData.ConfigWxData.Md5Key
// 	c := md5.New()
// 	_, err := c.Write([]byte(signStr))
// 	if err != nil {
// 		return "", err
// 	}
// 	signByte := c.Sum(nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.ToUpper(fmt.Sprintf("%x", signByte)), nil
// }

//检测配置
func (pc *NotifyClient) CheckConfig() {
	if pc.NotifyClientInterface.GetSignType() != "MD5" {
		pc.PrivateKey = util.Bytes2RSAPrivateKey([]byte(pc.ConfigData.ConfigAliData.RsaPrivateKey))
		pc.PublicKey = util.Bytes2RSAPublicKey([]byte(pc.ConfigData.ConfigAliData.AliPublicKey))
		if pc.ConfigData.ConfigAliData.RsaPrivateKey == nil {
			errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
		}
	} else {
		if pc.ConfigData.ConfigAliData.Key == "" {
			errors.ThrewError(errors.PAY_CONFIG_NO_KEY)
		}
	}

}

// GetSign 产生签名
func (pc *NotifyClient) GetSign(m map[string]string) (string, error) {
	delete(m, "sign")
	if pc.NotifyClientInterface.IsOldPay() {
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
	signData := strings.Join(data, "&")
	sign := ""
	switch pc.NotifyClientInterface.GetSignType() {
	case "MD5":
		sign = pc.Md5Sign(signData)
	case "RSA":
		sign = pc.RsaSign(signData)
	case "RSA2":
		sign = pc.Rsa2Sign(signData)
	}
	return sign, nil
}

func (pc *NotifyClient) Md5Sign(signData string) string {
	signData = signData + pc.ConfigData.ConfigAliData.Key
	c := md5.New()
	_, err := c.Write([]byte(signData))
	if err != nil {
		return ""
	}
	signByte := c.Sum(nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", signByte)
}

func (pc *NotifyClient) RsaSign(signData string) string {
	s := sha1.New()
	_, err := s.Write([]byte(signData))
	if err != nil {
		log.Println(err)
	}
	hashByte := s.Sum(nil)
	signByte, err := pc.PrivateKey.Sign(rand.Reader, hashByte, crypto.SHA1)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signByte)
}

func (pc *NotifyClient) Rsa2Sign(signData string) string {
	s := sha256.New()
	_, err := s.Write([]byte(signData))
	if err != nil {
		log.Println(err)
	}
	hashByte := s.Sum(nil)
	signByte, err := pc.PrivateKey.Sign(rand.Reader, hashByte, crypto.SHA256)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signByte)
}

func (pc *NotifyClient) Send() interface{} {
	return pc.NotifyClientInterface.BuildResData()
}

func (pc *NotifyClient) IsOldPay() bool {
	return false
}
