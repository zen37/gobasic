package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type person struct {
	XMLName   xml.Name `xml:"Person"`
	Age       int      `xml:"Age,attr"`
	FirstName string   `xml:"FirstName"`
	Lastname  string   `xml:"LastName"`
	Hobbies   []string `xml:"Hobbies>Hobby"`
}

/* type Hobby struct {
	//XMLName xml.Name `xml:"Hobby"`
	Value string `xml:"Hobby"`
	Type  string `xml:"Type,attr"`
} */

func main() {
	xmlFile, err := ioutil.ReadFile("file.xml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\tSuccessfully opened file.xml")
	//fmt.Println(string(xmlFile))
	fmt.Println()

	var p person
	if err := xml.Unmarshal(xmlFile, &p); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.FirstName, p)

	for _, v := range p.Hobbies {
		fmt.Println(v)
	}
}
