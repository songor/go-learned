// 使用 Write 方法向客户端发送响应
package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := "<html><head><title>Go Web Programming</title></head><body><h1>Hello World</h1></body></html>"
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/write", writeExample)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
