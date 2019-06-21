package main

import (
	"fmt"
	"github.com/jerryzhengj/gomirco/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"golang.org/x/net/context"
)

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *hello_world.HelloRequest, rsp *hello_world.HelloResponse) error {
	fmt.Printf("received request from %s", req.String())
	rsp.Greeting = "Hello World: " + req.Name
	return nil
}

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"http://192.168.10.110:2379",
		}
	})
	service := micro.NewService(
		micro.Name("hello_world"),
		micro.Version("1.0"),
		micro.Registry(reg),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	hello_world.RegisterHelloWorldHandler(service.Server(), new(HelloWorld))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
