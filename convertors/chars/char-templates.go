package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type StatSet struct {
	XMLName xml.Name `xml:"set"`

	Name string `xml:"name,attr"`
	Val  string `xml:"val,attr"`
}

type L2JCharTemplate struct {
	XMLName xml.Name `xml:"class"`

	Name      string `xml:"name,attr"`
	ClassId   uint32 `xml:"id,attr"`
	BaseLevel uint32 `xml:"baseLevel,attr"`

	Set []StatSet `xml:"set"`
}

type List struct {
	XMLName xml.Name          `xml:"list"`
	List    []L2JCharTemplate `xml:"class"`
}

func main() {
	fd, _ := os.Open("C:\\Users\\og_ki\\AppData\\Roaming\\JetBrains\\GoLand2021.2\\scratches\\scratch.xml")
	bytes, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	var list = &List{}
	xml.Unmarshal(bytes, list)

	for _, t := range list.List {
		log.Printf("%+v", t)
	}

	log.Printf("%+v", list)
}
