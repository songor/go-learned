// 通过编写首部实现客户端重定向
package main

import "net/http"

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/redirect", headerExample)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
