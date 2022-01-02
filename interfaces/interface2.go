package main

//http://fast4ward.online/posts/a-guide-to-interfaces-in-go/

import (
	"fmt"
)

type foobar interface {
	foo()
	bar()
}

type itemA struct{}

func (a *itemA) foo() {
	fmt.Println("foo on A")
}

func (a *itemA) bar() {
	fmt.Println("bar on A")
}

type itemB struct {
	name string
}

func (v *itemB) foo() {
	fmt.Println("foo on B")
}

func (v *itemB) bar() {
	fmt.Println("bar on B")
}

func (v itemB) bar2() {
	fmt.Println("bar on B")
}

func doFooA(item *itemA) {
	item.foo()
}

func doFoo(item foobar) {
	item.foo()
}

func main() {
	doFoo(&itemA{}) // Prints "foo on A"
	doFoo(&itemB{}) // Prints "foo on B"

	//doFoo(&itemB)   //does not compile - type itemB is not an expression
	/* does not compile - cannot use itemB{} (type itemB) as type foobar in argument to doFoo:
	   itemB does not implement foobar (bar method has pointer receiver)
	*/
	//doFoo(itemB{})
}
