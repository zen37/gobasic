package main

import (
	"fmt"
	"time"
)

var doneNumbers chan struct{}
var doneLetters chan struct{}

func main() {

	doneNumbers = make(chan struct{})
	doneLetters = make(chan struct{})

	go print_letters()
	go print_numbers()
	//	time.Sleep(5000 * time.Millisecond)

	<-doneLetters
	fmt.Println()
	fmt.Println("print_letters() finished")

	<-doneNumbers
	fmt.Println()
	fmt.Println("print_numbers() finished")
}

func print_letters() {

	fmt.Println("print_letters started")

	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf(" %c", i)
		time.Sleep(100 * time.Millisecond)
	}
	doneLetters <- struct{}{}
}

func print_numbers() {

	fmt.Println("print_numbers started")

	for i := 0; i <= 33; i++ {
		fmt.Print(" ", i)
		time.Sleep(100 * time.Millisecond)
	}
	doneNumbers <- struct{}{}
}
