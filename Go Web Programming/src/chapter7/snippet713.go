// 使用 Encoder 把结构编码为 JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func S713() {
	post := JsonPost{
		Id:      1,
		Content: "Hello World!",
		Author: JsonAuthor{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []JsonComment{
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

	jsonFile, err := os.Create("output.json")
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
