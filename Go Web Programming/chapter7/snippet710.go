// JSON 分析程序
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post710 struct {
	Id       int          `json:"id"`
	Content  string       `json:"content"`
	Author   Author710    `json:"author"`
	Comments []Comment710 `json:"comments"`
}

type Author710 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment710 struct {
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

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post Post710
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
