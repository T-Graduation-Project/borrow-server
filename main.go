package main

import (
	"github.com/T-Graduation-Project/borrow-server/app/api"
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/micro/go-micro/v2"
)

var (
	log = g.Log()
)

func main() {
	server := micro.NewService(
		micro.Name("borrow"),
		micro.Version("latest"),
	)
	server.Init()
	protobuf.RegisterBorrowHandler(server.Server(), new(api.BorrowApi))
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
