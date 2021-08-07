package main

/*
& Operator
& goes in front of a variable when you want to get that variable's memory address.

* Operator
* goes in front of a variable that holds a memory address and resolves it;
it is therefore the counterpart to the & operator.
It goes and gets the thing that the pointer was pointing at, e.g. *myString.
*/

import "fmt"

func main() {
	var x = 10
	fmt.Println("var x = 10 / x =", x)
	fmt.Println("&x =", &x)
	fmt.Println("*&x =", *&x)
	fmt.Println()
	var y *int
	fmt.Println("y *int / y =", y)
	y = &x
	fmt.Println("y = &x / y =", y)
	fmt.Println("&y =", &y)
	fmt.Println("*y =", *y)
	fmt.Println("*&y =", *&y)

}
