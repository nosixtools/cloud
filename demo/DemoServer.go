package main

import (
	"cloud/config"
	"time"
	"cloud/server"
	"cloud/demo/gen-go/hello/rpc"
)

type HelloServiceImpl struct {
	
}

func (this *HelloServiceImpl) Hello(name string ) (string, error)  {
	return "Hello:" + name, nil;
}

func main() {
	configTemp := &config.Config{Host:"127.0.0.1",Port:8889,Group:"test",Version:"1.0.0", Weight:"1"}
	registerConfig := &config.RegisterConfig{[]string{"localhost:2181"}, time.Second * 30}
	providerConfig := &config.ProviderConfig{Config:configTemp,RegisterConfig:registerConfig, ServiceName:"HelloService"}
	//url := &url.Url{Protocol:"cloud", ServiceName:"HelloService", Config:configTemp}
	//fmt.Println(url.GetUrl())
	handler := &HelloServiceImpl{}
	processor := rpc.NewHelloServiceProcessor(handler)
	server := server.NewThrifthServer(providerConfig, processor)
	server.Start()
	time.Sleep(time.Second * 1)
	for {
		if(!server.IsStarted()) {
			break
		}
	}
}
