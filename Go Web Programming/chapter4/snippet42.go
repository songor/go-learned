// 读取请求首部
package main

import (
	"fmt"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	log.Println(r.Header["Accept-Encoding"])
	log.Println(r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, h)
}

func main() {
	http.HandleFunc("/headers", headers)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
