package client

import (
	"cloud/config"
	"git.apache.org/thrift.git/lib/go/thrift"
	"reflect"
	"strconv"
	"time"
	"errors"
)

var (
	ErrSerAddress          = errors.New("cloud: service address error")
	waitNodeDelaySecond    = time.Second * 1
	waitNodeDelay          = 1
	ErrMethodNotExists     = errors.New("cloud: method not exists")
	ErrProxyExec           = errors.New("cloud: params error")
	ErrEmptyHosts          = errors.New("cloud: empty hosts")
	ErrServicesUnavaliable = errors.New("cloud: services unavaliable")
)

type ThriftClient struct {
	referenceConfig *config.ReferenceConfig
	service interface{}
	inited bool
}

func NewThriftClient(config *config.ReferenceConfig, service interface{}) *ThriftClient {
	return &ThriftClient{config, service, false}
}

func (tc *ThriftClient) Exec(methodName string, params ...interface{}) ([]reflect.Value, error) {
	if tc.service == nil {
		return nil, nil
	}

	if !tc.inited {
		tc.createClientFactory()
	}

	proxy := reflect.ValueOf(tc.service)
	exec := proxy.MethodByName(methodName)
	if !exec.IsValid() {
		return nil, ErrMethodNotExists
	}
	param := make([]reflect.Value, len(params))
	for i, item := range params {
		param[i] = reflect.ValueOf(item)
	}
	return exec.Call(param), nil
}

func (tc *ThriftClient) createClientFactory() error {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	address := tc.referenceConfig.Config.Host + ":" + strconv.Itoa(tc.referenceConfig.Config.Port)
	if address == "" {
		return ErrEmptyHosts
	}
	transport, err := thrift.NewTSocket(address)
	if err != nil {
		return  err
	}
	useTransport := transportFactory.GetTransport(transport)
	client := reflect.ValueOf(tc.service)
	mutable := client.Elem()
	mutable.FieldByName("Transport").Set(reflect.Value(reflect.ValueOf(useTransport)))
	mutable.FieldByName("ProtocolFactory").Set(reflect.Value(reflect.ValueOf(protocolFactory)))
	mutable.FieldByName("InputProtocol").Set(reflect.Value(reflect.ValueOf(protocolFactory.GetProtocol(useTransport))))
	mutable.FieldByName("OutputProtocol").Set(reflect.Value(reflect.ValueOf(protocolFactory.GetProtocol(useTransport))))
	mutable.FieldByName("SeqId").SetInt(0)
	if err := transport.Open(); err != nil {
		return  err
	}
	tc.inited = true
	return nil
}
