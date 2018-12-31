package common

import "github.com/openpeng/fool-pay/constant"

type ConfigWxData struct {
	AppId       string   `json:"app_id"`  // 公众账号ID
	MchId       string   `json:"mch_id"`  // 商户id
	Md5Key      string   `json:"md5_key"` // md5 秘钥
	AppCertPem  string   `json:"app_cert_pem"`
	AppKeyPem   string   `json:"app_key_pem"`
	SignType    string   `json:"sign_type"` // MD5  HMAC-SHA256
	LimitPay    []string `json:"limit_pay"` // 指定不能使用信用卡支付   不传入，则均可使用
	FeeType     string   `json:"fee_type"`  // 货币类型  当前仅支持该字段
	NotifyUrl   string   `json:"notify_url"`
	RedirectUrl string   `json:"redirect_url"` // 如果是h5支付，可以设置该值，返回到指定页面
}

type ConfigAliData struct {
	UseSandbox bool `json:"use_sandbox"` // 是否使用沙盒模式
	OldMd5     bool `json:"old_md5"`     //老版本还保留了MD5的加密方式
	//安全检验码
	Key string `json:"key"`
	//签约支付宝账号
	SellerEmail string `json:"seller_email"`
	//合作身份者id
	Partner  string `json:"partner"`
	AppId    string `json:"app_id"`
	SignType string `json:"sign_type"` // RSA  RSA2
	// 可以填写文件路径，或者密钥字符串  当前字符串是 rsa2 的支付宝公钥
	AliPublicKey string `json:"ali_public_key"`
	// 可以填写文件路径，或者密钥字符串  我的沙箱模式，rsa与rsa2的私钥相同，为了方便测试
	RsaPrivateKey string   `json:"rsa_private_key"`
	LimitPay      []string `json:"limit_pay"` // 指定不能使用信用卡支付   不传入，则均可使用
	// 		'limit_pay' => [
	// 			//'balance',// 余额
	// 			//'moneyFund',// 余额宝
	// 			//'debitCardExpress',//     借记卡快捷
	// //        'creditCard',//信用卡
	// 			//'creditCardExpress',// 信用卡快捷
	// 			//'creditCardCartoon',//信用卡卡通
	// 			//'credit_group',// 信用支付类型（包含信用卡卡通、信用卡快捷、花呗、花呗分期）
	// 		],// 用户不可用指定渠道支付当有多个渠道时用“,”分隔
	// 与业务相关参数
	NotifyUrl string `json:"notify_url"`
	ReturnUrl string `json:"return_url"`
}
type ConfigPubData struct {
	ReturnRaw bool `json:"return_raw"` // 在处理回调时，是否直接返回原始数据，默认为false
}
type BaseConfig struct {
	ConfigWxData   ConfigWxData
	ConfigAliData  ConfigAliData
	ConfigPubData  ConfigPubData
	PayThirdServer constant.PayServer
}
