package common

type SignClientInterface interface {
	BuildSignData(signData map[string]string) string
	Sign(signData map[string]string) string
}

type SignWay string

const (
	SIGN_WAY_WX_MD5   SignWay = "sign_way_wx_md5"
	SIGN_WAY_ALI_MD5  SignWay = "sign_way_ali_md5"
	SIGN_WAY_ALI_RSA  SignWay = "sign_way_ali_rsa"
	SIGN_WAY_ALI_RSA2 SignWay = "sign_way_ali_rsa2"
)

type SignHandle struct {
	SignClientInterface
	SignType SignWay
}

func NewSignClient(hci SignClientInterface) *SignHandle {
	return &SignHandle{
		SignClientInterface: hci,
	}
}
