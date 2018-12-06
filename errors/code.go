package errors

type ErrorCode int

const (
	HANDLE_OK  ErrorCode = 0
	PAY_MODULE ErrorCode = 10001 + iota
	PAY_PUB_WRONG
	NO_SUPPORT_CHANNEL
)

var ErrorMessage = map[ErrorCode]string{
	HANDLE_OK:          "处理成功",
	PAY_MODULE:         "支付内部错误",
	NO_SUPPORT_CHANNEL: "不支持的支付类型",
}

type PayError struct {
	ErrorCode ErrorCode
	Message   string
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
		}
	}
}
