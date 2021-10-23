// 使用处理器函数处理请求
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
