// 通过 HTTPS 提供服务
package main

import (
	"log"
	"net/http"
)

func S34() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
