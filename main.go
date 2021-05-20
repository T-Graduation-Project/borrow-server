package main

import (
	"github.com/T-Graduation-Project/borrow-server/app/api"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

var (
	log = g.Log()
)

func main() {
	borrowApi := &api.BorrowApi{}
	borrowApi.InitClient()

	server := micro.NewService(
		micro.Name("borrow"),
		micro.Registry(etcdv3.NewRegistry(
			registry.Addrs(g.Cfg().GetString("registry_addr")),
		)),
	)
	server.Init()
	protobuf.RegisterBorrowHandler(server.Server(), borrowApi)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
