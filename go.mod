module github.com/T-Graduation-Project/borrow-server

go 1.15

require (
	github.com/fullstorydev/grpcurl v1.8.0 // indirect
	github.com/gogf/gf v1.15.1
	github.com/gogf/katyusha v0.0.0-20201219155529-faa901bdd86f
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.4.3
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0
)

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)
