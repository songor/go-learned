// 创建模板自定义函数
package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func processFunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmp-func.html").Funcs(funcMap)
	t, _ = t.ParseFiles("chapter5/tmp-func.html")
	t.Execute(w, time.Now())
}

func main() {
	http.HandleFunc("/process-func", processFunc)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
