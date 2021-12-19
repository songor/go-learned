package main

import "encoding/xml"

type XmlPost struct {
	XMLName xml.Name  `xml:"post"`
	Id      string    `xml:"id,attr"`
	Content string    `xml:"content"`
	Author  XmlAuthor `xml:"author"`
	Xml     string    `xml:",innerxml"`
}

type XmlAuthor struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type XmlComment struct {
	Id      string    `xml:"id,attr"`
	Content string    `xml:"content"`
	Author  XmlAuthor `xml:"author"`
}

type JsonPost struct {
	Id       int           `json:"id"`
	Content  string        `json:"content"`
	Author   JsonAuthor    `json:"author"`
	Comments []JsonComment `json:"comments"`
}

type JsonAuthor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JsonComment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
