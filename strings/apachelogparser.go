package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	log         string = "apache.log"
	log_request string = "apache_requests.log"
	// StandardEnglishFormat is the time layout to use for parsing date/time format
	StandardEnglishFormat = "02/Jan/2006:15:04:05 -0700"
	dt                    = "[10/Oct/2000:13:55:36 -1100]"
)

func main() {

	//read_access_log()
	//read_requests()

	date, err := readDateTime(dt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(date)

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

func readDateTime(input string) (d time.Time, err error) {
	if input[0] != '[' {
		err = fmt.Errorf("got %q, want '['", input[0])
		return
	}
	idx := strings.Index(input, "]")
	if idx == -1 {
		err = errors.New("missing closing ']'")
		return
	}
	if d, err = time.Parse(StandardEnglishFormat, input[1:idx]); err != nil {
		err = errors.New("failed to parse datetime: " + err.Error())
	}
	return
}
