//https://go.dev/play/p/trb2sE3MXj_v

package main

import (
	"fmt"
	"regexp"
)

//"/[a-f0-9]+/css/"
//const match = "/[a-f0-9]+/images/"
const match = "/[a-f0-9]+/css/"

//const path = "/35e77b/images/map-icons/catering.png HTTP/1.1"
const path = "/7b0744/css/vegguide-combined.css HTTP/1.1"

var re = regexp.MustCompile(match)

func main() {

	str := path // "/save//kjjsad"

	fmt.Printf("Match: %v - Request: %v\n", re.String(), path) // print pattern
	fmt.Println("does request match pattern?", re.MatchString(str))
}
