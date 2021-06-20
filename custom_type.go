package main

import "fmt"

type anyName string

func (a anyName) print() {
	fmt.Println("method print() called, anyName is now: ", a)
}

func (a anyName) copy() {
	//does not change the receiver
	fmt.Println("[func (a anyName) pointer()] called")
	fmt.Println("anyName receiver is: ", a)
	fmt.Println("anyName receiver gets a new value: \"hallo welt\"")
	a = "hallo welt"
	fmt.Println("a is: ", a)
	fmt.Println("&a is: ", &a)
}

func (a *anyName) pointer() {
	//changes the receiver
	fmt.Println("[func (a *anyName) pointer()] called")
	fmt.Println("anyName receiver is: ", a)
	fmt.Println("*anyName receiver is: ", *a)
	fmt.Println("anyName receiver gets a new value: \"hallo welt\"")
	*a = "hallo welt"
	fmt.Println("a is: ", a)
	fmt.Println("*a is: ", *a)
	fmt.Println("&a is: ", &a)
	fmt.Println("*&a is: ", *&a)
}

func main() {
	var x anyName

	x = "hello world"
	fmt.Println("x = \"hello world\"")
	fmt.Println()

	x.copy()
	x.print()

	fmt.Println()

	x.pointer()
	x.print()

}
