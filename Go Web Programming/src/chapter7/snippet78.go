// 手动将 Go 结构编码至 XML
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func S78() {
	post := XmlPost{
		Id:      "1",
		Content: "Hello World!",
		Author: XmlAuthor{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	xmlFile, err := os.Create("output.xml")
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
