package main

import (
	"context"
	"fmt"
	"log"

	"github.com/99n/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/99n/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("gomall-consul-1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Responce res: %v\n", res)
}
