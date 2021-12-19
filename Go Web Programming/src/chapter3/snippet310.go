// 串联两个处理器函数
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func logHandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func S310() {
	http.HandleFunc("/hello", logHandlerFunc(hello))
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
