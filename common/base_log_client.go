package common

type LogClientInterface interface {
	Write(message interface{})
}

type LogClient struct {
	LogClientInterface
}

func NewLogClient(hci LogClientInterface) *LogClient {
	return &LogClient{
		LogClientInterface: hci,
	}
}
