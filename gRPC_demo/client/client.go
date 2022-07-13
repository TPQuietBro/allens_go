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
	conn, err := grpc.Dial(":1990", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("连接失败")
	}

	defer conn.Close()

	client := service.NewProdServiceClient(conn)

	stock, err1 := client.GetProductStock(context.Background(), &service.ProductRequest{
		ProId: 1990,
	})

	if err1 != nil {
		log.Fatalln("库存查询失败")
	}

	fmt.Println("库存为: ", stock.ProStock)
}
