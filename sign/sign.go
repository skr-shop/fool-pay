package sign

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
)

type Sign struct {
}

// GetSign 产生签名
func (pc *Sign) GetSign(m map[string]string) (string, error) {
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

func (pc *Sign) Md5Sign(signData string) string {
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

func (pc *Sign) RsaSign(signData string) string {
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

func (pc *Sign) Rsa2Sign(signData string) string {
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
func (pc *Sign) CheckSign(data string, sign string) error {
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

func (pc *Sign) Send() interface{} {
	return pc.ClientInterface.BuildResData()
}

func (pc *Sign) IsOldPay() bool {
	return false
}
