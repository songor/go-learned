// 对 XML 进行分析
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func S72() {
	xmlFile, err := os.Open("post.xml")
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

	var post XmlPost
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
