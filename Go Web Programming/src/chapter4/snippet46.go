// 通过 MultipartForm 字段接收用户上传的文件
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func processFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["upload"][0]
	file, err := fileHeader.Open()

	//file, _, err := r.FormFile("upload")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func S46() {
	http.HandleFunc("/process-file", processFile)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
