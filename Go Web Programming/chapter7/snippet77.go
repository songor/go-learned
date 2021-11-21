// 使用 Marshal 函数生成 XML 文件
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post77 struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author77 `xml:"author"`
}

type Author77 struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post77{
		Id:      "1",
		Content: "Hello World!",
		Author: Author77{
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
	err = ioutil.WriteFile("chapter7/output.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
