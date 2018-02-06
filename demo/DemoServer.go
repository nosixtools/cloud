package main

import (
	//"cloud/demo/gen-go/hello/rpc"
	"cloud/config"
	//"cloud/server"
	"time"
	//"cloud/common/url"
	//"fmt"
	//"cloud/register"
	log "code.google.com/p/log4go"

	"cloud/register"
)

type HelloServiceImpl struct {
	
}

func (this *HelloServiceImpl) Hello(name string ) (string, error)  {
	return "Hello:" + name, nil;
}

func main() {
	configTemp := &config.Config{Host:"127.0.0.1",Port:8888,Group:"test",Version:"1.0.0", Weight:"1"}
	registerConfig := &config.RegisterConfig{[]string{"192.168.1.17:4180"}, time.Second * 30}
	providerConfig := &config.ProviderConfig{Config:configTemp,RegisterConfig:registerConfig, ServiceName:"HelloService"}
	//url := &url.Url{Protocol:"cloud", ServiceName:"HelloService", Config:configTemp}
	//fmt.Println(url.GetUrl())
	//handler := &HelloServiceImpl{}
	//processor := rpc.NewHelloServiceProcessor(handler)
	//server := server.NewThrifthServer(providerConfig, processor)
	//server.Start()
	//time.Sleep(time.Second * 1)
	register, err := register.NewRegister(registerConfig)
	if err != nil {
		log.Error(err.Error())
	}
	register.RegisterService(providerConfig)
	//for {
	//	if(!server.IsStarted()) {
	//		break
	//	}
	//}
}
