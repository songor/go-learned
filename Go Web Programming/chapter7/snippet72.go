// 对 XML 进行分析
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post72 struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author72    `xml:"author"`
	Xml      string    `xml:",innerxml"`
}

type Author72 struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("chapter7/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post72
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
