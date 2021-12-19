// 迭代动作
package main

import (
	"html/template"
	"net/http"
)

func processRange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp-range.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func S55() {
	http.HandleFunc("/process-range", processRange)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
