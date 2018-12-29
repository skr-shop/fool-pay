package alinotify

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"github.com/openpeng/fool-pay/common"
	alibase "github.com/openpeng/fool-pay/common/ali"
)

type ClientInterface interface {
	BuildResData() string
}

type NotifyClient struct {
	*alibase.NotifyClient
	ClientInterface
}

func NewNotifyClient(configData common.BaseConfig, intface interface{}) *NotifyClient {
	var tmp = &NotifyClient{
		ClientInterface: intface.(ClientInterface),
	}
	tmp.NotifyClient = alibase.NewNotifyClient(configData, intface.(alibase.NotifyClientInterface))
	return tmp
}

func (pc *NotifyClient) GetSignType() string {
	switch strings.ToUpper("RSA") {
	case "RSA2":
		return "RSA2"
	case "RSA":
		return "RSA"
	}
	return "RSA"
}

// CheckSign 检测签名
func (ac *NotifyClient) CheckSign(signData, sign string) error {
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	s := sha1.New()
	_, err = s.Write([]byte(signData))
	if err != nil {
		return err
	}
	hash := s.Sum(nil)
	return rsa.VerifyPKCS1v15(ac.PublicKey, crypto.SHA1, hash, signByte)
}
