# borrow-server
借阅服务
## HOW TO RUN
~~~bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
./borrow-server --gf.gcfg.file=config.toml

docker build -t csighub.tencentyun.com/urftian/tendoc-oms-server-go .
docker run --rm -it -p 8199:8199 csighub.tencentyun.com/urftian/tendoc-oms-server-go
~~~

## 开发常用命令
~~~bash
# proto 协议生成代码
protoc --proto_path=./proto --go_out=. borrow.proto

# model 默认使用 config.toml的[database.default]
gf gen dao
gf gen dao -l "mysql:username:password@tcp(ip:port)/tablename"

# docker
docker rmi csighub.tencentyun.com/urftian/tendoc-oms-server-go
docker build -t csighub.tencentyun.com/urftian/tendoc-oms-server-go .
docker run --rm -it -p 58080:58080 csighub.tencentyun.com/urftian/tendoc-oms-server-go

docker login 
docker push csighub.tencentyun.com/urftian/tendoc-oms-server-go
~~~