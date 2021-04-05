module github.com/T-Graduation-Project/borrow-server

go 1.15

require (
	github.com/T-Graduation-Project/book-server v0.0.0-20210324071718-a7c51745ddcb
	github.com/T-Graduation-Project/proto v1.0.0 // indirect
	github.com/T-Graduation-Project/user-server v0.0.0-20210331081427-d241087f0205
	github.com/gogf/gf v1.15.4
	github.com/golang/protobuf v1.5.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/examples v0.0.0-20210324172016-702608ffae4d // indirect
)

replace google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.26.0