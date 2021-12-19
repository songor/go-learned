// 带有附加配置的 Web 服务器
package main

import "net/http"

func S32() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
