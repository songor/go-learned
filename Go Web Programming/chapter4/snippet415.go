// 从请求的首部获取 cookie
package main

import (
	"fmt"
	"net/http"
)

func getCookie(w http.ResponseWriter, r *http.Request) {
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)

	// 使用 Cookie 方法和 Cookies 方法
	c, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	fmt.Fprintln(w, c)

	cs := r.Cookies()
	fmt.Fprintln(w, cs)
}

func main() {
	http.HandleFunc("/get-cookie", getCookie)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
