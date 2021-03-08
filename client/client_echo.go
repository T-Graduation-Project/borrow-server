package main

import (
	"github.com/T-Graduation-Project/borrow-server/protobuf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/katyusha/krpc"
	"golang.org/x/net/context"
	"time"
)

func main() {
	conn, err := krpc.Client.NewGrpcClientConn("borrow")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	echoClient := protobuf.NewBorrowClient(conn)
	for i := 0; i < 500; i++ {
		res, err := echoClient.GetBookList(
			context.Background(), &protobuf.GetBookListReq{})
		if err != nil {
			g.Log().Error(err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		g.Log().Print("Response:", res)
	}
}
