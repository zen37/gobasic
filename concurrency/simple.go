package main

import (
	"fmt"
	"time"
)

func main() {
	go print_letters()
	go print_numbers()
	time.Sleep(5000 * time.Millisecond)
}

func print_letters() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf(" %c", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func print_numbers() {
	for i := 0; i <= 27; i++ {
		fmt.Print(" ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
