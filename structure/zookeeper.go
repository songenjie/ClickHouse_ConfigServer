package structure

import "encoding/xml"

type Zookeeper struct {
	XMLName xml.Name `xml:"yandex"`
	Text    string   `xml:",chardata"`
	Logger  struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level"`
		Size  string `xml:"size"`
		Count string `xml:"count"`
	} `xml:"logger"`
	Zookeeper struct {
		Text string  `xml:",chardata"`
		Node []*Node `xml:"node"`
	} `xml:"zookeeper"`
}

type Node struct {
	Text  string `xml:",chardata"`
	Index string `xml:"index,attr"`
	Host  string `xml:"host"`
	Port  int    `xml:"port"`
}
