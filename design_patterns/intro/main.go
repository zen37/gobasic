package main

import "fmt"

type duck interface {
	//quack()
	//swim()
	display()
}

type mallard struct {
	name      string
	gender    string
	birthyear int
}

func main() {

	var m mallard
	m = mallard{"cute", "m", 2020}

	fmt.Println(m)

	var d duck
	d = mallard{"cute", "m", 2020}
	d.display()
}

func (p mallard) display() {
	fmt.Println("I'm a mallard")
}
