package main

/*
to do, punctuation marks should not be considered part of the word
*/

import (
	"fmt"
	"strings"
)

type txt struct {
	word  string
	count int
}

// const text = "the dog is friends with another dog."
const text = "betty bought the butter, the butter was bitter, betty bought more butter to make the bitter butter better "

func main() {
	count(text)
}

func count(input string) {

	s := strings.Fields(input)
	fmt.Println(s)

	m := make(map[string]int, len(s))

	for _, v := range s {
		m[v]++
	}

	fmt.Println(m)
}
