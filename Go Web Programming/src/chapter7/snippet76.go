// 使用 Decoder 分析 XML
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func S76() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment XmlComment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
