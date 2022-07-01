package service

import (
	"context"
	"distributedemo/registraty"
	"fmt"
	"log"
	"net/http"
)

// 集中启动服务
func Start(ctx context.Context, host, port string, reg registraty.Registration, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	// 注册服务
	err := registraty.RegisterService(reg)

	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registraty.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	// 处理服务启动异常情况
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	// 手动终止服务
	go func() {
		fmt.Printf("%v started, press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
