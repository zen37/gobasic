package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	file, err := ioutil.ReadFile("xml/file.xml")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", file)

	fmt.Println("---------------------------------")

	text := string(file)
	fmt.Println(text)
}
