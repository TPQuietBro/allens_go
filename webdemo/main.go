package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main()  {
	// 利用gin框架解决cors 跨域问题
	g := gin.Default()
	g.Use(AllowAll())
	g.GET("/home/getuser", func(ctx *gin.Context) {
		res := make(map[string]string)
		res["name"] = "allen"
		json, err := json.Marshal(res)
		if err != nil {
			log.Fatal("invalid json")
		}

		ctx.String(200, string(json))
	})
	_ = g.Run(":8080")

	// http.HandleFunc("/home", handleLogin)
	// http.HandleFunc("/home/getuser", handleGetUser)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func AllowAll() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}
	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.Form
	fmt.Println(params);
	for k, v := range params {
		tempV := strings.Join(v, "");
		fmt.Printf("k = %v, v = %v \n", k, tempV);
		value, err := strconv.Atoi(tempV);
		if err != nil {
			log.Fatal("invalid value")
		}
		if value == 1001 {
			res := make(map[string]string)
			res["name"] = "allen"
			json, err := json.Marshal(res)
			if err != nil {
				log.Fatal("invalid json")
			}
			fmt.Println(string(json))
			// fmt.Fprintln(w, string(json))
			w.Write([]byte(string(json)))
		}
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello my go"))
	r.ParseForm()
	params := r.Form
	fmt.Println(params);
	for k, v := range params {
		fmt.Printf("k = %v, v = %v \n", k, strings.Join(v, ""));
	}

	fmt.Fprintln(w, "hello my go")
	// fmt.Println(r.URL.Query())
}