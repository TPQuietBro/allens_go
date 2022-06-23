package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/home", handleReq)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello my go"))
}