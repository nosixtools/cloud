package main

import (
	"cloud/client"
	"cloud/config"
	"cloud/demo/gen-go/hello/rpc"
	"fmt"
)

func main() {
	configTemp := &config.Config{Host:"127.0.0.1",Port:8888,Group:"test",Version:"1.0.0", Weight:"1"}
	referenceConfig := &config.ReferenceConfig{configTemp}
	helloService := &rpc.HelloServiceClient{}

	client := client.NewThriftClient(referenceConfig, helloService)
	fmt.Println(client.Exec("Hello", "nosix"))
	fmt.Println(client.Exec("Hello", "nosix"))
}
