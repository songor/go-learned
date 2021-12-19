// 为了展示模板中的上下文感知特性而设置的处理器
package main

import (
	"html/template"
	"net/http"
)

func processContextAware(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp-context-aware.html")
	context := `I asked: <i>"What's up?"</i>`
	t.Execute(w, context)
}

func S517() {
	http.HandleFunc("/process-context-aware", processContextAware)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
