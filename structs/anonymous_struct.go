package main

/*
from the book Learning Go by Jon Bodner
a variable can implement a struct type without first giving the struct type a name.
this is called an anonymous struct.
why it’s useful to have a data type that’s only associated with a single instance?
	1. when you translate external data into a struct or a struct into external data (like JSON or protocol buffers).
	this is called unmarshaling and marshaling data
	2. writing tests
*/

import "fmt"

func main() {

	var person struct {
		name    string
		dob     string
		country string
	}

	person.name = "bob"
	person.dob = "19700101"
	person.country = "uk"

	fmt.Println(person)

	pet := struct {
		name string
		age  int
		kind string
	}{
		name: "fido",
		age:  3,
		kind: "dog",
	}

	fmt.Println(pet)

}
