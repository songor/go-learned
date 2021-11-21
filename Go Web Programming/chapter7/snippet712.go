// 将结构封装为 JSON
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post712 struct {
	Id       int          `json:"id"`
	Content  string       `json:"content"`
	Author   Author712    `json:"author"`
	Comments []Comment712 `json:"comments"`
}

type Author712 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment712 struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post712{
		Id:      1,
		Content: "Hello World!",
		Author: Author712{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment712{
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

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("chapter7/output.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
