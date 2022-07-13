package main

import (
	"fmt"
	"gRPC_demo/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	rpc := grpc.NewServer()

	service.RegisterProdServiceServer(rpc, service.Pro)

	listen, err := net.Listen("tcp", ":1990")
	if err != nil {
		log.Fatalln("监听失败")
	}
	err1 := rpc.Serve(listen)

	if err1 != nil {
		log.Fatalln("启动失败")
	}
	fmt.Println("服务启动成功")
}
