package main

import (
	"log"
	"net"

	"github.com/oldwind/grpc_demo/server/services"
	"google.golang.org/grpc"
)

func main() {

	// just test
	rpcServer := grpc.NewServer()
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
