package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func registerHomeRoute() {
	http.HandleFunc("/home", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"code": 1,
		"data": "首页",
	}
	listJson, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintln(w, string(listJson))
}
