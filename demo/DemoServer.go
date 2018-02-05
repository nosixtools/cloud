package main

import (
	"cloud/demo/gen-go/hello/rpc"
	"cloud/config"
	"cloud/server"
	"time"
	"cloud/common/url"
	"fmt"
)

type HelloServiceImpl struct {
	
}

func (this *HelloServiceImpl) Hello(name string ) (string, error)  {
	return "Hello:" + name, nil;
}

func main() {
	configTemp := &config.Config{Host:"127.0.0.1",Port:8888,Group:"test",Version:"1.0.0", Weight:"1"}
	url := &url.Url{Protocol:"cloud", ServiceName:"HelloService", Config:configTemp}
	fmt.Println(url.GetUrl())
	providerConfig := &config.ProviderConfig{Config:configTemp}
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
