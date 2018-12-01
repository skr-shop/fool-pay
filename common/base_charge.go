package common

import (
	"fmt"
)

type ChargeClientInterface interface {
	Send() interface{}
	SetSign() interface{}
}

type ChargeClient struct {
	ConfigData BaseConfig
	ChargeClientInterface
}

func NewChargeClient(intface ChargeClientInterface) ChargeClient {
	return ChargeClient{
		ChargeClientInterface: intface,
	}
}

type PubChargeClient struct {
	This       PubChargeClientImpl
	ConfigData BaseConfig
	PubChargeClientImpl
}

type PubChargeClientImpl interface {
	ChargeClientInterface
}

func (pc PubChargeClient) Send() interface{} {
	fmt.Println(pc.This.SetSign())
	return "公共的发送"
}

func (pc PubChargeClient) SetSign() interface{} {
	return "签名有没有?"
}
