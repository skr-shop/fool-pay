package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/imroc/req"
)

type HttpClient struct {
	Debug       bool
	HttpRequest *http.Request
}

func InitHttpClient() *HttpClient {
	var httpClient = new(HttpClient)
	httpClient.HttpRequest = new(http.Request)
	req.SetTimeout(5 * time.Second)
	return httpClient
}

func (hc *HttpClient) handle(paramData map[string]interface{}, headerParam map[string]string) (req.Param, req.Header) {
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if headerParam != nil {
		for k, v := range headerParam {
			header[k] = v
		}
	}
	if hc.Debug {
		req.Debug = true
	}
	param := req.Param(paramData)
	return param, headerParam
}

func (hc *HttpClient) Post(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Post(url, param, header)
	var data interface{}
	if err != nil {
		fmt.Println("http:request:"+url+"===error", err)
	} else {
		r.ToJSON(&data)
	}
	return data, err
}

func (hc *HttpClient) Put(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Put(url, param, header)
	var data interface{}
	if err != nil {
		fmt.Println("http:request:"+url+"===error", err)
	} else {
		r.ToJSON(&data)
	}
	return data, err
}

func (hc *HttpClient) Delete(url string, paramData map[string]interface{}, headerParam map[string]string) ([]byte, error) {
	param, header := hc.handle(paramData, headerParam)
	r, _ := req.Delete(url, param, header)
	return r.ToBytes()
}

func (hc *HttpClient) Get(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Get(url, param, header)
	var data interface{}
	if err != nil {
		fmt.Println("http:request:"+url+"===error", err)
	} else {
		r.ToJSON(&data)
	}
	return data, err
}

func (hc *HttpClient) PostBodyJson(url string, body interface{}) (interface{}, error) {
	var data interface{}
	r, err := req.Post(url, req.BodyJSON(&body))
	if err != nil {
		fmt.Println("http:request:"+url+"===error", err)
	} else {
		r.ToJSON(&data)
	}
	return data, err
}

func (hc *HttpClient) PostBodyJsonWithHeader(url string, body interface{}) (interface{}, error) {
	r, err := req.Post(url, req.BodyJSON(&body), req.Header{"origin": "tiku.gaodun.coms"})
	var data interface{}
	if err != nil {
		fmt.Println("http:request:"+url+"===error", err)
	} else {
		r.ToJSON(&data)
	}
	return data, err
}

func (hc *HttpClient) PostBodyXml(url string, body interface{}) ([]byte, error) {
	r, err := req.Post(url, body)
	return r.Bytes(), err
}
