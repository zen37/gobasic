package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var x interface{} = "abc"
	convertInterfaceToString(x)

	var y interface{} = []int{1, 2, 3}
	convertInterfaceToString(y)

	var z interface{} = person{"joe", 20}
	convertInterfaceToString(z)

}

func convertInterfaceToString(i interface{}) {

	str := fmt.Sprintf("%v", i)
	fmt.Println(str)

}
