package common

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"log"

	"github.com/openpeng/fool-pay/constant"
	"github.com/openpeng/fool-pay/util"
)

type SignClientInterface interface {
	BuildSignData() string
}

type SignKey struct {
	PrivateKey string
	PublicKey  string
}
type SignConfig struct {
	SignType constant.SignWay
	SignKey
}
type SignHandle struct {
	SignClientInterface
	SignConfig
}

func NewSignClient(sc SignConfig, sci SignClientInterface) *SignHandle {
	return &SignHandle{
		SignConfig:          sc,
		SignClientInterface: sci,
	}
}

func (sh *SignHandle) Sign() (string, error) {

	signData := sh.SignClientInterface.BuildSignData()
	sign := ""
	switch sh.SignConfig.SignType {
	case "MD5":
		sign = sh.Md5Sign(signData)
	case "RSA":
		sign = sh.RsaSign(signData)
	case "RSA2":
		sign = sh.Rsa2Sign(signData)
	}
	return sign, nil
}

func (pc *SignHandle) Md5Sign(signData string) string {
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

func (pc *SignHandle) RsaSign(signData string) string {

	s := sha1.New()
	_, err := s.Write([]byte(signData))
	if err != nil {
		log.Println(err)
	}
	hashByte := s.Sum(nil)
	signByte, err := util.Bytes2RSAPrivateKey([]byte(pc.PrivateKey)).Sign(rand.Reader, hashByte, crypto.SHA1)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signByte)
}

func (pc *SignHandle) Rsa2Sign(signData string) string {
	s := sha256.New()
	_, err := s.Write([]byte(signData))
	if err != nil {
		log.Println(err)
	}
	hashByte := s.Sum(nil)
	signByte, err := util.Bytes2RSAPrivateKey([]byte(pc.PrivateKey)).Sign(rand.Reader, hashByte, crypto.SHA256)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signByte)
}

// CheckSign 检测签名
func (pc *SignHandle) CheckSign(sign string) error {
	signData := pc.BuildSignData()
	var s hash.Hash
	var whichHash crypto.Hash
	fmt.Println(pc.SignType)
	switch pc.SignType {
	case "MD5":
		fmt.Println(signData)
		if sign == pc.Md5Sign(signData) {
			return nil
		}
		return errors.New("签名验证失败")
	case "RSA":
		s = sha1.New()
		whichHash = crypto.SHA1
	case "RSA2":
		s = sha256.New()
		whichHash = crypto.SHA256
	default:
		return errors.New("不支持的加密方式")
	}

	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	_, err = s.Write([]byte(signData))
	if err != nil {
		return err
	}
	hashByte := s.Sum(nil)
	return rsa.VerifyPKCS1v15(util.Bytes2RSAPublicKey([]byte(pc.PublicKey)), whichHash, hashByte[:], signByte)
}
