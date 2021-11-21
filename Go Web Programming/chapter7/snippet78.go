// 手动将 Go 结构编码至 XML
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post78 struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author78 `xml:"author"`
}

type Author78 struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post78{
		Id:      "1",
		Content: "Hello World!",
		Author: Author78{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	xmlFile, err := os.Create("chapter7/output.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}

	xmlFile.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
