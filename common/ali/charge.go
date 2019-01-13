package ali

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/openpeng/fool-pay/util"

	"github.com/openpeng/fool-pay/container"

	"github.com/openpeng/fool-pay/errors"

	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/common/ali/data"
)

type ClientInterface interface {
	BuildData() string
	GetSignType() string
	CheckConfig()
	IsOldPay() bool
	BuildResData() interface{}
}

// 调用方法的时候要指定是ClientInterface的方法，
// 这样才能实现调用到子集的数据
// 如果自身也实现了就可以实现不会报错的问题（子集继承上一层会继承到方法，实现会自动关联）
type ChargeClient struct {
	*common.ChargeClient
	*common.HttpClient
	ClientInterface
	*data.AliResResult
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewChargeClient(configData common.BaseConfig, intface interface{}) *ChargeClient {
	var cc = &ChargeClient{
		ClientInterface: intface.(ClientInterface),
		HttpClient:      container.HttpClient,
		AliResResult:    &data.AliResResult{},
	}
	//将继承的实现掉 配置信息和请求都是在公共的地方，可以考虑先统一实现再初始化配置，考虑先处理配置是可以判断配置有没有问题
	cc.ChargeClient = common.NewChargeClient(configData, intface.(common.ChargeClientInterface))
	return cc
}

//对外的入口，接入请求数据
func (pc *ChargeClient) Charge(data common.ReqData) interface{} {
	pc.ChargeClient.ReqData = data
	pc.ClientInterface.CheckConfig() //可以调用到子集的方法
	pc.ClientInterface.BuildData()   //这么调用是因为本身没实现，必定走的是子集的方法,但是还是推荐加ClientInterface
	return pc.Send()
}

//检测配置
func (pc *ChargeClient) CheckConfig() {
	if pc.ClientInterface.GetSignType() != "MD5" {
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

// GetSign 产生签名
func (pc *ChargeClient) GetSign(m map[string]string) (string, error) {
	delete(m, "sign")
	if pc.ClientInterface.IsOldPay() {
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
	switch pc.ClientInterface.GetSignType() {
	case "MD5":
		sign = pc.Md5Sign(signData)
	case "RSA":
		sign = pc.RsaSign(signData)
	case "RSA2":
		sign = pc.Rsa2Sign(signData)
	}
	return sign, nil
}

func (pc *ChargeClient) Md5Sign(signData string) string {
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

func (pc *ChargeClient) RsaSign(signData string) string {
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

func (pc *ChargeClient) Rsa2Sign(signData string) string {
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

// CheckSign 检测签名
func (pc *ChargeClient) CheckSign(data string, sign string) error {
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	s := sha1.New()
	var whichHash crypto.Hash = crypto.SHA1
	switch pc.ClientInterface.GetSignType() {
	case "RSA2":
		s = sha256.New()
		whichHash = crypto.SHA256
	}
	_, err = s.Write([]byte(data))
	if err != nil {
		return err
	}
	hash := s.Sum(nil)
	return rsa.VerifyPKCS1v15(pc.PublicKey, whichHash, hash[:], signByte)
}

func (pc *ChargeClient) Send() interface{} {
	return pc.ClientInterface.BuildResData()
}

func (pc *ChargeClient) IsOldPay() bool {
	return false
}
