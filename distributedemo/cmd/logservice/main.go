package main

import (
	"context"
	"distributedemo/log"
	"distributedemo/registraty"
	"distributedemo/service"
	"fmt"
	stlog "log"
)

func main() {
	// 初始化自定义log
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	address := fmt.Sprintf("%s:%s", host, port)

	reg := registraty.Registration{
		ServiceName: "Log Service",
		ServiceUrl:  address,
	}

	// 开启日志服务
	ctx, err := service.Start(context.Background(), host, port, reg, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}
	// 等待服务关闭
	<-ctx.Done()
	fmt.Println("shutting down log service")
}
