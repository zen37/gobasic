//https://ibraheem.ca/writings/go-in-twenty/

package main

import "fmt"

type aStruct struct {
	s   string
	b   bool
	i   int
	i64 int64
	f   float64
}

type number struct {
	value int
}

func main() {
	//create new struct with a struct literals
	a := aStruct{s: "hello world", b: true}
	fmt.Println("aStruct{s: \"hello world\", b: true} =>", a) // -> {hello world true 0 0 0}

	a = aStruct{b: true, i64: 2, f: 0.89637623478}
	fmt.Println(a) // -> { true 0 2 0.89637623478}

	//pointer to a struct
	var ap *aStruct
	ap = &aStruct{}
	fmt.Println("&aStruct{} =>", ap) //-> &{ false 0 0 0}

	ap = &aStruct{"hallo welt", true, 2, 1234567890, 0.78787}
	fmt.Println(ap) //-> &{hallo welt true 2 1234567890 0.78787}

	fmt.Println("&ap.b =>", &ap.b)
	fmt.Println("ap.b =>", ap.b)
	fmt.Println("*&ap.b =>", *&ap.b)

	n := number{value: 20}
	fmt.Printf("[method n.isPositive] is the number [ %v ] positive?, %v\n", n.value, n.isPositive())
	fmt.Printf("[function isPositive(n)] is the number [ %v ] positive?, %v\n", n.value, isPositive(n))

	n = number{value: -10}
	fmt.Printf("is the number [ %v ] positive?, %v\n", n.value, n.isPositive())

}
func (n number) isPositive() bool {
	return n.value >= 0
}

func isPositive(n number) bool {
	return n.value >= 0
}
