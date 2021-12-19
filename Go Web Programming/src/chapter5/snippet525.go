// 使用显式定义的模板
package main

import (
	"html/template"
	"net/http"
)

func processLayout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func S525() {
	http.HandleFunc("/process-layout", processLayout)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
