// 对表单进行语法分析
package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	http.HandleFunc("/process", process)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
