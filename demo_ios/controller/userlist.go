package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StringMap map[string]interface{}
type SringMapList []StringMap

func registerUserListRoute() {
	http.HandleFunc("/home/getlist", listHandler)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	list := SringMapList{}
	keyName := "name"
	keyAge := "age"
	keyDec := "description"
	data1 := StringMap{
		keyName: "唐鹏",
		keyAge:  18,
		keyDec:  "成熟又稳重",
	}

	data2 := StringMap{
		keyName: "marry",
		keyAge:  21,
		keyDec:  "pretty and kindly",
	}

	data3 := StringMap{
		keyName: "金州拉文",
		keyAge:  35,
		keyDec:  "三分贼准",
	}
	list = append(list, data1, data2, data3)

	result := map[string]interface{}{
		"code": 1,
		"data": list,
	}

	listJson, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err.Error())
	}
	jsonString := string(listJson)
	fmt.Println(jsonString)
	fmt.Fprintln(w, jsonString)
}
