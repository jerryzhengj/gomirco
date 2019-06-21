package main

import (
	"fmt"
	"github.com/jerryzhengj/gomirco/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"golang.org/x/net/context"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.10.110:2379",//本机ip是192.168.10.110
		}
	})
	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	greeter := hello_world.NewHelloWorldService("hello_world", service.Client()) // 创建服务hello_world的client对象, 以便调用其中定义的RPC方法'Hello'

	rsp, err := greeter.Hello(context.TODO(), &hello_world.HelloRequest{Name: "go"}) // 传入HelloWorldRequest对象作为调用RPC方法的参数'Hello'
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
