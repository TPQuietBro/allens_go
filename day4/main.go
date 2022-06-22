package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start := time.Now();
	ch := make(chan string)
	
	someApis := []string{
		"https://baidu.com",
		"https://qq.com",
		"https://youtube.com",
		"https://ddddddd123.com",
	}

	for _, api := range someApis {
		go checkApi(api, ch)
	}
	for i:=0; i < len(someApis); i++ {
		res := <-ch
		log(res)
	}
	end := time.Since(start);
	log("cost: " + fmt.Sprintf("%f",end.Seconds()))
}

func checkApi(api string, ch chan string)  {
	_, err := http.Get(api)
	if !hasError(err, ch) {
		// log(api + ", done")
		ch<-fmt.Sprintf("%s done", api)
	} 
}

func hasError(err error, ch chan string) bool {
	if err != nil {
		// log(err.Error())
		ch <- err.Error()
		return true
	}
	return false
}

func log(info string) {
	fmt.Println(info)
}