// 串联多个处理器
package main

import (
	"fmt"
	"net/http"
)

type HiHandler struct{}

func (h HiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}

func logHandler(h http.Handler) http.Handler {
	// 把 匿名函数 转换成了 HandlerFunc 类型
	// 该类型实现了 ServeHTTP 接口
	// 所以其也可以转换成 Handler 类型
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {
	hi := HiHandler{}
	http.Handle("/hi", logHandler(hi))
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
