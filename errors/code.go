package errors

import "github.com/openpeng/fool-pay/container"

type ErrorCode int

const (
	HANDLE_OK  ErrorCode = 0
	PAY_MODULE ErrorCode = 10000 + iota
	PAY_PUB_WRONG
	NO_SUPPORT_CHANNEL
	SIGN_WRONG
	PAY_CONFIG_NO_KEY
	PAY_DATA_TRS_ERROR
	PAY_WAY_NO_SIGN_TYPE
)

var ErrorMessage = map[ErrorCode]string{
	HANDLE_OK:            "处理成功",
	PAY_MODULE:           "支付内部错误",
	NO_SUPPORT_CHANNEL:   "不支持的支付类型",
	SIGN_WRONG:           "签名错误",
	PAY_CONFIG_NO_KEY:    "缺少加密的key",
	PAY_DATA_TRS_ERROR:   "数据转失败",
	PAY_WAY_NO_SIGN_TYPE: "加密方式和支付类型不匹配",
}

type PayError struct {
	ErrorCode ErrorCode `json:"error_code"`
	Message   string    `json:"message"`
}

func ThrewError(errorCode ErrorCode) {
	if mes, has := ErrorMessage[errorCode]; has {
		panic(PayError{ErrorCode: errorCode, Message: mes})
	}
	panic(PayError{ErrorCode: PAY_MODULE, Message: ErrorMessage[PAY_MODULE]})
}

func ThrewMessageError(message string) {
	panic(PayError{ErrorCode: PAY_PUB_WRONG, Message: message})
}

func IsOK() PayError {
	return PayError{ErrorCode: HANDLE_OK, Message: ErrorMessage[HANDLE_OK]}
}

// Catch ...
func Catch(e *PayError) {
	if err := recover(); err != nil {
		switch err.(type) {
		case PayError:
			pe := err.(PayError)
			e.ErrorCode = pe.ErrorCode
			e.Message = pe.Message
			return
		}
		container.LogHandle.Write(err)
	}
}
