package main

/*readme
from https://ibraheem.ca/writings/go-in-twenty/
- an anonymous function is useful when you want to define a function inline without having to name it
- an anonynous function allows to dynamically change what the function does at runtime
- a function can take functions as arguments
- a function can also return an anonymous function
- a closure is an anonymous function that can reference a variable from outside itself
*/

import (
	"fmt"
)

//a function can take functions as arguments
func doSomething(funcParam func() string) {
	fmt.Println(funcParam())

}

//a function can also return an anonymous function
func returnAnoFunc() func() string {
	return func() string {
		return "a function can also return an anonymous function"
	}
}

func main() {

	//an anonymous function is useful when you want to define a function inline
	//without having to name it
	ano := func() string {
		return "define a function inline without having to name it"
	}
	fmt.Println("ano(): ", ano())

	//an anonynous function allows to dynamically change what the func. does at runtime
	ano = func() string {
		return "re-define the same anonymous function"
	}
	fmt.Println("ano(): ", ano())

	doSomething(func() string {
		return "a function can take functions as arguments"
	})

	fmt.Println("returnAnoFunc():", returnAnoFunc())
	r := returnAnoFunc()
	fmt.Println("r := returnAnoFunc():", r())

	//anonymous functions can form closures.
	//a closure is a an anonymous function that can reference
	//a variable out of itself

	x := 42
	closure := func() {
		x += 1 //x = x + 1
		fmt.Println("x in closure: ", x)
	}
	closure()
	closure()
}
