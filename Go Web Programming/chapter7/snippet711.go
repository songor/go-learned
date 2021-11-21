// 使用 Decoder 对 JSON 进行语言分析
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post711 struct {
	Id       int          `json:"id"`
	Content  string       `json:"content"`
	Author   Author711    `json:"author"`
	Comments []Comment711 `json:"comments"`
}

type Author711 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment711 struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("chapter7/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post711
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}
