// JSON 分析程序
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func S710() {
	jsonFile, err := os.Open("post.json")
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

	var post JsonPost
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
