package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	log         string = "apache.log"
	log_request string = "apache_requests.log"
)

func main() {

	read_access_log()
	//read_requests()

}

func read_access_log() {

	record := 0

	//f, err := os.Open("apache.log")
	f, err := os.Open(log)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		record++

		s := strings.Fields(string(scanner.Bytes()))
		//print request
		fmt.Println(record, s[5][1:], s[6])
	}

}

func read_requests() {

	//var mapCounter map[string]int

	mapCounter := make(map[string]int)

	records := 0

	f, err := os.Open(log_request)
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		records++

		s := strings.Fields(string(scanner.Bytes()))
		http_method := s[1]
		mapCounter[http_method]++

	}
	fmt.Println("records=", records)
	fmt.Println(mapCounter)
}
