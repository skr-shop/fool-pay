package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net"
	"time"
)

//RandomStr 获取一个随机字符串
func RandomStr() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

func MapStringToStruct(m map[string]string, i interface{}) error {
	bin, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bin, i)
	if err != nil {
		return err
	}
	return nil
}

func Bytes2RSAPrivateKey(priKey []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priKey)
	if block == nil {
		fmt.Println("Sign private key decode error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return privateKey.(*rsa.PrivateKey)
}

// 阿里官网的测试
// var privateKey = []byte(`
// -----BEGIN RSA PRIVATE KEY-----
// MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
// 7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
// Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
// AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
// ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
// XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
// /jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
// IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
// 4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
// DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
// 9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
// DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
// AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
// -----END RSA PRIVATE KEY----- `)
// var publicKey = []byte(`
// -----BEGIN PUBLIC KEY-----
// MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
// ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
// wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
// AUeJ6PeW+DAkmJWF6QIDAQAB
// -----END PUBLIC KEY----- `)

// func RsaDecrypt(ciphertext []byte) ([]byte, error) {
// 	block, _ := pem.Decode(privateKey)
// 	if block == nil {
// 		return nil, errors.New("private key error!")
// 	}
// 	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
// }
// func RsaEncrypt(origData []byte) ([]byte, error) {
// 	block, errs := pem.Decode(publicKey)
// 	if block == nil {
// 		fmt.Println(string(errs))
// 		return nil, errors.New("public key error")
// 	}
// 	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pub := pubInterface.(*rsa.PublicKey)
// 	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
// }

func Bytes2RSAPublicKey(pubKey []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		fmt.Println("Sign pubilc key decode error")
	}
	pubilcKey, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
	}
	return pubilcKey.(*rsa.PublicKey)
}
