package common

import "github.com/openpeng/fool-pay/constant"

// type BaseDataInterface interface {
// 	CheckDataParam()
// 	BuildReqData(reqd ReqData)
// }

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
	UseSandbox string `json:"use_sandbox"` // 是否使用沙盒模式
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
	RsaPrivateKey string `json:"rsa_private_key"`
	LimitPay      string `json:"limit_pay"` // 指定不能使用信用卡支付   不传入，则均可使用
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

// <?php
// /**
//  * @author: helei
//  * @createTime: 2016-07-28 18:05
//  * @description: 支付相关接口的数据基类
//  */

// namespace Payment\Common;

// use Payment\Utils\ArrayUtil;

// /**
//  * Class BaseData
//  * 支付相关接口的数据基类
//  * @package Payment\Common\Weixin\Dataa
//  */
// abstract class BaseData
// {

//     /**
//      * 支付的请求数据
//      * @var array $data
//      */
//     protected $data;

//     /**
//      * 支付返回的数据
//      * @var array $retData
//      */
//     protected $retData;

//     /**
//      * BaseData constructor.
//      * @param ConfigInterface $config
//      * @param array $reqData
//      * @throws PayException
//      */
//     public function __construct(ConfigInterface $config, array $reqData)
//     {
//         $this->data = array_merge($config->toArray(), $reqData);

//         try {
//             $this->checkDataParam();
//         } catch (PayException $e) {
//             throw $e;
//         }
//     }

//     /**
//      * 获取变量，通过魔术方法
//      * @param string $name
//      * @return null|string
//      * @author helei
//      */
//     public function __get($name)
//     {
//         if (isset($this->data[$name])) {
//             return $this->data[$name];
//         }

//         return null;
//     }

//     /**
//      * 设置变量
//      * @param $name
//      * @param $value
//      * @author helei
//      */
//     public function __set($name, $value)
//     {
//         $this->data[$name] = $value;
//     }

//     /**
//      * 设置签名
//      * @author helei
//      */
//     public function setSign()
//     {
//         $this->buildData();

//         $values = ArrayUtil::removeKeys($this->retData, ['sign']);

//         $values = ArrayUtil::arraySort($values);

//         $signStr = ArrayUtil::createLinkstring($values);

//         $this->retData['sign'] = $this->makeSign($signStr);
//     }

//     /**
//      * 返回处理之后的数据
//      * @return array
//      * @author helei
//      */
//     public function getData()
//     {
//         return $this->retData;
//     }

//     /**
//      * 签名算法实现  便于后期扩展微信不同的加密方式
//      * @param string $signStr
//      * @return string
//      */
//     abstract protected function makeSign($signStr);

//     /**
//      * 构建用于支付的签名相关数据
//      * @return array
//      */
//     abstract protected function buildData();

//     /**
//      * 检查传入的参数. $reqData是否正确.
//      * @return mixed
//      * @throws PayException
//      */
//     abstract protected function checkDataParam();
// }
