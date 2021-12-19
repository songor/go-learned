package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	// 创建一个用于运行测试的多路复用器
	mux = http.NewServeMux()
	// 绑定想要测试的处理器
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/3", nil)
	// 向被测试的处理器发送请求
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 3 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/3", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
