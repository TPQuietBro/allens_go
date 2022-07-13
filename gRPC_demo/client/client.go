package main

import (
	"context"
	"fmt"
	"gRPC_demo/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 客户端新起一个连接, 指定服务端端口, 并设置证书(防止报错)
	conn, err := grpc.Dial(":1990", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("连接失败")
	}
	// 适当时候关闭连接
	defer conn.Close()
	// pb业务绑定这个链接
	client := service.NewProdServiceClient(conn)
	// 像调用本地服务一样调用远程服务, 并传入参数
	stock, err1 := client.GetProductStock(context.Background(), &service.ProductRequest{
		ProId: 1990,
	})

	if err1 != nil {
		log.Fatalln("库存查询失败")
	}

	fmt.Println("库存为: ", stock.ProStock)
}
