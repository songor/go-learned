// 使用 Encoder 把结构编码为 JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post713 struct {
	Id       int          `json:"id"`
	Content  string       `json:"content"`
	Author   Author713    `json:"author"`
	Comments []Comment713 `json:"comments"`
}

type Author713 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment713 struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post713{
		Id:      1,
		Content: "Hello World!",
		Author: Author713{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment713{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("chapter7/output.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
