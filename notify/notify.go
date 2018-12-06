package notify

type NotifyProcessData struct {
	GetNotifyData
}

type Notify interface {
	GetNotifyData()
	NotifyProcess(data NotifyProcessData)
}
