// 使用 Marshal 函数生成 XML 文件
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func S77() {
	post := XmlPost{
		Id:      "1",
		Content: "Hello World!",
		Author: XmlAuthor{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	//output, err := xml.Marshal(&post)
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("output.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
