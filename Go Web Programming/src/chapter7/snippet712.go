// 将结构封装为 JSON
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func S712() {
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

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("output.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
