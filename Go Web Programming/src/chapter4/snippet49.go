// 通过 WriteHeader 方法将状态码写入到响应当中
package main

import (
	"fmt"
	"net/http"
)

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func S49() {
	http.HandleFunc("/write-header", writeHeaderExample)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
