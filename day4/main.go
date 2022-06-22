package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start := time.Now();

	someApis := []string{
		"https://baidu.com",
		"https://qq.com",
		"https://youtube.com",
		"https://ddddddd123.com",
	}

	for _, api := range someApis {
		go checkApi(api)
	}

	end := time.Since(start);
	log("cost: " + fmt.Sprintf("%f",end.Seconds()))

	time.Sleep(time.Second * 3)
}

func checkApi(api string)  {
	_, err := http.Get(api)
	if !hasError(err) {
		log(api + ", done")
	} 
}

func hasError(err error) bool {
	if err != nil {
		log(err.Error())
		return true
	}
	return false
}

func log(info string) {
	fmt.Println(info)
}