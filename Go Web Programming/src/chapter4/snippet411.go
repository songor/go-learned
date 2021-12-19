// 编写 JSON 输出
package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User:    "Sau Sheong",
		Threads: []string{"first", "second", "third"},
	}
	body, _ := json.Marshal(post)
	w.Write(body)
}

func S411() {
	http.HandleFunc("/json", jsonExample)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
