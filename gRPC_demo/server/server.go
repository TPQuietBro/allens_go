package main

import (
	"fmt"
	"gRPC_demo/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 新起一个grpc的服务
	rpc := grpc.NewServer()
	// pb注册grpc的服务
	// 并绑定实现: service.Pro 实现了业务方法`GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)`
	service.RegisterProdServiceServer(rpc, service.Pro)
	// 开启监听
	listen, err := net.Listen("tcp", ":1990")
	if err != nil {
		log.Fatalln("监听失败")
	}
	// grpc关联这个监听, 并启动服务
	err1 := rpc.Serve(listen)

	if err1 != nil {
		log.Fatalln("启动失败")
	}
	fmt.Println("服务启动成功")
}
