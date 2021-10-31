// 设置动作
package main

import (
	"html/template"
	"net/http"
)

func processWith(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter5/tmp-with.html")
	t.Execute(w, "hello")
}

func main() {
	http.HandleFunc("/process-with", processWith)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
