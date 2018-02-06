package server

import (
	"cloud/config"
	"git.apache.org/thrift.git/lib/go/thrift"
	"strconv"
	"fmt"
	"cloud/register"
	log "code.google.com/p/log4go"
)

type ThriftServer struct {
	providerConfig *config.ProviderConfig
	processor thrift.TProcessor
	server *thrift.TSimpleServer
	hasStarted bool
}

func NewThrifthServer(config *config.ProviderConfig, processor thrift.TProcessor)  *ThriftServer  {
	server  := &ThriftServer{providerConfig:config, processor:processor, hasStarted:false}
	return server
}

func(ts *ThriftServer) Start() error  {
	var err error
	//start service
	go func(err error)  {
		NetworkAddr := ts.providerConfig.Config.Host+":"+strconv.Itoa(ts.providerConfig.Config.Port)
		transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
		if err != nil {
			fmt.Println(err.Error())
		}
		ts.server = thrift.NewTSimpleServer4(ts.processor, serverTransport, transportFactory, protocolFactory)
		fmt.Println("thrift server in", NetworkAddr)
		ts.hasStarted = true;
		err = ts.server.Serve()
	}(err)
	if err == nil {
		register, err := register.NewRegister(ts.providerConfig.RegisterConfig)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		err = register.RegisterService(ts.providerConfig)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}
	return err
}

func (ts *ThriftServer) Stop() error  {
	return ts.server.Stop()
}

func (ts *ThriftServer) IsStarted() bool {
	return ts.hasStarted
}


