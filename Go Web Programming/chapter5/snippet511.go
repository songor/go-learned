// 调用嵌套模板的处理器
package main

import (
	"html/template"
	"net/http"
)

func processTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter5/t1.html", "chapter5/t2.html")
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/process-template", processTemplate)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
