// 使用 gob 包读写二进制数据
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type GobPost struct {
	Id      int
	Content string
	Author  string
}

func storePost(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func loadPost(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := GobPost{
		Id:      1,
		Content: "Hello World!",
		Author:  "Sau Sheong",
	}
	storePost(post, "post")
	var postRead GobPost
	loadPost(&postRead, "post")
	fmt.Println(postRead)
}
