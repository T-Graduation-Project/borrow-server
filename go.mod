module github.com/T-Graduation-Project/borrow-server

go 1.15

require (
	github.com/gogf/gf v1.15.1
	github.com/gogf/katyusha v0.0.0-20201219155529-faa901bdd86f
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)
