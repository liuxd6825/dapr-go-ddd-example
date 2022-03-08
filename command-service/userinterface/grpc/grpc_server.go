package grpc

import (
	svr "github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/grpc/service"
	pd "github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/grpc/service/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Start() *grpc.Server {
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()
	// 2. 将刚刚我们新建的ProdService注册进去
	pd.RegisterProdServiceServer(rpcServer, new(svr.ProdService))
	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
	return rpcServer
}
