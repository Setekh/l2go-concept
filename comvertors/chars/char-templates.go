package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type StatSet struct {
	XMLName xml.Name `xml:"set"`

	Name string `xml:"name,attr"`
	Val  string `xml:"val,attr"`
}

type L2JCharTemplate struct {
	XMLName xml.Name `xml:"class"`

	Name    string `xml:"name,attr"`
	ClassId uint32 `xml:"class_id,attr"`
	Level   uint32 `xml:"level,attr"`

	Set []StatSet `xml:"set"`
}

type List struct {
	XMLName xml.Name          `xml:"list"`
	List    []L2JCharTemplate `xml:"class"`
}

func main() {
	fd, _ := os.Open("/Users/vladcazacu/Library/Application Support/JetBrains/GoLand2021.2/scratches/scratch_1.xml")
	bytes, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	var list = &List{}
	xml.Unmarshal(bytes, list)

}
