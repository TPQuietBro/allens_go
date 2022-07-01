package main

import (
	"context"
	"distributedemo/registraty"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registraty.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registraty.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("registraty services started, press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("shutting down registraty service")
}
