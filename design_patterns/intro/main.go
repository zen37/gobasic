package main

import "fmt"

type duck interface {
	//quack()
	//swim()
	identify()
}

type mallard struct {
	name      string
	gender    string
	birthyear int
}

type redhead struct {
}

func main() {

	var m mallard
	m = mallard{"cute", "m", 2020}

	fmt.Println(m)

	var d duck
	d = mallard{"cute", "m", 2020}
	//d.identify()
	whoAreYou(d)

	d = redhead{}
	//d.identify()
	whoAreYou(d)
}

func (p mallard) identify() {
	fmt.Println("I'm a mallard")
}

func (p redhead) identify() {
	fmt.Println(("I'm a redhead duck"))
}

func whoAreYou(d duck) {
	d.identify()
}
