// 在处理器里面生成一个随机数
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func processIf(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp-if.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func S53() {
	http.HandleFunc("/process-if", processIf)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
