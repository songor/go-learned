// 使用 HttpRouter 实现的服务器
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func helloWithParams(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", helloWithParams)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
