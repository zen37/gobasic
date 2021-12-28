package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const url = "https://jsonplaceholder.typicode.com"

//const url = "http://httpbin.org/get"

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*100))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
		//panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		//panic(err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	s := string(b)

	fmt.Println(s)
}
