package main

import (
	"github.com/T-Graduation-Project/borrow-server/app/api"
	"github.com/T-Graduation-Project/borrow-server/app/service"
	borrow_server "github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/katyusha/krpc"
)

func main() {
	c := krpc.Server.NewGrpcServerConfig()
	c.Options = append(
		c.Options,
		krpc.Server.ChainUnary(
			service.Interceptor.UnaryCtx,
			service.Interceptor.UnaryValidate,
		),
	)
	s := krpc.Server.NewGrpcServer(c)
	borrow_server.RegisterEchoServer(s.Server, api.Echo)
	borrow_server.RegisterUserServer(s.Server, api.User)
	s.Run()
}