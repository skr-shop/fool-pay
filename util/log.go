package util

import "fmt"

type LogClient struct {
}

func InitLogClient() *LogClient {
	return &LogClient{}
}

//记录日志
func (lc *LogClient) Write(message interface{}) {
	fmt.Println(message)
}
