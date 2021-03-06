// 使用多个处理器对请求进行处理
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func S37() {
	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
