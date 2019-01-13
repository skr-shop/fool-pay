package container

import (
	"github.com/openpeng/fool-pay/common"
	"github.com/openpeng/fool-pay/util"
)

var (
	HttpClient *common.HttpClient
	LogHandle  *common.LogClient
)

func init() {
	HttpClient = common.NewHttpClient(util.InitHttpClient())
	LogHandle = common.NewLogClient(util.InitLogClient())
}
