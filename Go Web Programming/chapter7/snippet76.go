// 使用 Decoder 分析 XML
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Author76 struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment76 struct {
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author76 `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("chapter7/post.xml")
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
				var comment Comment76
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
