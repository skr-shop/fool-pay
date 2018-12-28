package constant

// const (
// 	ALI_WEB = iota + 1
// 	ALI_APP
// 	WECHAT
const VERSION = "0.0.1"

type PayChannel string
type PayOperation string
type PayStatus string
type PayServer string
type AliPayMethod string

const (
	PAY_SERVER_ALI PayServer = "ali"

	PAY_SERVER_WX PayServer = "wx"

	// 支付相关常量
	//========================= ali相关接口 =======================//
	ALI_CHANNEL_APP PayChannel = "ali_app" // 支付宝 手机app 支付

	ALI_CHANNEL_WAP PayChannel = "ali_wap" // 支付宝 手机网页 支付

	ALI_CHANNEL_WEB PayChannel = "ali_web" // 支付宝 PC 网页支付

	ALI_CHANNEL_QR PayChannel = "ali_qr" // 支付宝 扫码支付

	ALI_CHANNEL_BAR PayChannel = "ali_bar" // 支付宝 条码支付

	ALI_CHANNEL_OLD_QUICK PayChannel = "ali_quick" // 支付宝 老版本快捷MD5/RAS加密 支付

	//========================= 微信相关接口 =======================//
	// 支付常量
	WX_CHANNEL_APP PayChannel = "wx_app" // 微信 APP 支付

	WX_CHANNEL_PUB PayChannel = "wx_pub" // 微信 公众账号 支付

	WX_CHANNEL_QR PayChannel = "wx_qr" // 微信 扫码支付  (可以使用app的帐号，也可以用公众的帐号完成)

	WX_CHANNEL_BAR PayChannel = "wx_bar" // 微信 刷卡支付，与支付宝的条码支付对应

	WX_CHANNEL_LITE PayChannel = "wx_lite" // 微信小程序支付

	WX_CHANNEL_WAP PayChannel = "wx_wap" // 微信wap支付，针对特定用户

	// 其他操作常量
	ALI_CHARGE PayOperation = "ali_charge" // 支付

	ALI_REFUND PayOperation = "ali_refund" // 退款

	ALI_RED PayOperation = "ali_red" // 红包

	ALI_TRANSFER PayOperation = "ali_transfer" // 转账

	// 其他相关常量
	WX_CHARGE PayOperation = "wx_charge" // 支付

	WX_REFUND PayOperation = "wx_refund" // 退款

	WX_RED PayOperation = "wx_red" // 红包

	WX_TRANSFER PayOperation = "wx_transfer" // 转账

	//======================= 交易状态常量定义 ======================//
	TRADE_STATUS_SUCC PayStatus = "success" // 交易成功

	TRADE_STATUS_FAILD PayStatus = "not_pay" // 交易未完成

	//========================= 金额问题设置 =======================//
	PAY_MIN_FEE = "0.01" // 支付的最小金额

	PAY_MAX_FEE = "100000000.00" // 支付的最大金额

	TRANS_FEE = "50000" // 转账达到这个金额，需要添加额外信息

	//========================= 支付宝各类method名称 =======================//
	// wap 支付
	WAP_PAY_METHOD AliPayMethod = "alipay.trade.wap.pay"

	// app 支付
	APP_PAY_METHOD AliPayMethod = "alipay.trade.app.pay"

	// 即时到账 web支付
	PC_PAY_METHOD AliPayMethod = "create_direct_pay_by_user"

	// 扫码支付   用户扫商户的二维码
	QR_PAY_METHOD AliPayMethod = "alipay.trade.precreate"

	// 条码支付   商户扫用户的二维码
	BAR_PAY_METHOD AliPayMethod = "alipay.trade.pay"

	// 统一收单线下交易查询
	TRADE_QUERY_METHOD AliPayMethod = "alipay.trade.query"

	// 统一收单交易退款查询  未完成
	REFUND_QUERY_METHOD AliPayMethod = "alipay.trade.fastpay.refund.query"

	// 转账情况查询
	TRANS_QUERY_METHOD AliPayMethod = "alipay.fund.trans.order.query"

	// 统一收单交易退款接口
	TRADE_REFUND_METHOD AliPayMethod = "alipay.trade.refund"

	// 单笔转账到支付宝账户接口
	TRANS_TOACCOUNT_METHOD AliPayMethod = "alipay.fund.trans.toaccount.transfer"
)
