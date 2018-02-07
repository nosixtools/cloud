package main

import (
	"cloud/client"
	"cloud/demo/gen-go/hello/rpc"
	"cloud/config"
	"fmt"
	"time"
	//"cloud/register"
)

func main() {
	configTemp := &config.Config{Group:"test",Version:"1.0.0", Weight:"1"}
	registerConfig := &config.RegisterConfig{[]string{"localhost:2181"}, time.Second * 30}
	referenceConfig := &config.ReferenceConfig{configTemp, registerConfig, "HelloService"}

	//register, _ := register.NewRegister(referenceConfig.RegisterConfig)
	//register.DiscoverServices(referenceConfig)
	//fmt.Println(register.Seletor())
	helloService := &rpc.HelloServiceClient{}
	//
	client := client.NewThriftClient(referenceConfig, helloService)
	fmt.Println(client.Exec("Hello", "nosix"))
	fmt.Println(client.Exec("Hello", "nosix"))
}
