package common

type HttpClientInterface interface {
	Post(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error)
	Put(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error)
	Get(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error)
	Delete(url string, paramData map[string]interface{}, headerParam map[string]string) ([]byte, error)
	PostBodyJson(url string, body interface{}) (interface{}, error)
	PostBodyXml(url string, body interface{}) ([]byte, error)
}

type HttpClient struct {
	HttpClientInterface
}

func NewHttpClient(hci HttpClientInterface) *HttpClient {
	return &HttpClient{
		HttpClientInterface: hci,
	}
}
