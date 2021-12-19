// 处理器使用在不同模板文件中定义的同名模板
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func processLayoutSameTemplate(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	random := rand.Intn(10)
	if random >= 1 && random < 3 {
		t, _ = template.ParseFiles("layout-same-template.html", "red_hello.html")
	} else if random >= 3 && random < 6 {
		t, _ = template.ParseFiles("layout-same-template.html", "blue_hello.html")
	} else {
		// 通过块动作定义默认模板
		t, _ = template.ParseFiles("layout-same-template.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func S528() {
	http.HandleFunc("/process-layout-same-template", processLayoutSameTemplate)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
