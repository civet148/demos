package main

import (
	"context"
	"github.com/civet148/demos/micro-grpc/serverpb"
	"time"

	"github.com/civet148/gotools/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/service/grpc"
)

/*
安装protoc
1.github上下载一个cpp包：https://github.com/google/protobuf/releases 后执行 make && make install安装即可
2.protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
3.安装protoc-gen-micro
go get github.com/micro/protoc-gen-micro
*/

const (
	ETCD_NAME_ECHO = "echo-server"
)

var regEtcd registry.Registry

type ServerImpl struct {
}

func init() {
	regEtcd = etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
}

func main() {

	go StartServer()
	StartClient()
	time.Sleep(10 * time.Second)
}

func StartServer() (err error) {
	service := grpc.NewService(
		micro.Registry(regEtcd),
		micro.RegisterInterval(10*time.Second),
		micro.RegisterTTL(5*time.Second),
		micro.Name(ETCD_NAME_ECHO),
		micro.Address("127.0.0.1:8877"),
	)

	if err = serverpb.RegisterEchoServerHandler(service.Server(), new(ServerImpl)); err != nil {
		log.Error(err.Error())
		return
	}
	if err = service.Run(); err != nil {
		log.Error()
		return
	}
	log.Infof("server over")
	return
}

func StartClient() {
	time.Sleep(1 * time.Second)
	client := grpc.NewService(
		micro.Registry(regEtcd),
	).Client()

	service := serverpb.NewEchoServerService(ETCD_NAME_ECHO, client)
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "lory",
		"X-From-Id": "10086",
	})
	for i := 0; i < 3; i++ {

		if pong, err := service.Call(ctx, &serverpb.Ping{Text: "Ping"}); err != nil {
			log.Error(err.Error())
			return
		} else {
			log.Infof("server response [%+v]", pong)
		}
		time.Sleep(1 * time.Second)
	}
	log.Infof("client over")
}

func (s *ServerImpl) Call(ctx context.Context, ping *serverpb.Ping, pong *serverpb.Pong) (err error) {
	md, _ := metadata.FromContext(ctx)
	log.Infof("md [%+v] req [%+v]", md, ping)
	pong.Text = "Pong"
	return
}
