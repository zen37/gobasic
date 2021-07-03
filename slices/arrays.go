package main

import "fmt"

func main() {

	// the array type is defined by its size + type of elements
	// by default an array is zero valued for ints is  0
	var a1 [3]int
	print("array [3]int", a1)

	var a2 = [3]int{1, 2, 3}
	fmt.Println("array [3]int{1, 2, 3} -> ", a2)

	//sparse array, most elements are set to their 0 values, can specify only the indices with values
	var a3 = [7]int{1, 4: 7, 8}
	print("array [7]int{1, 4: 7, 8}", a3)

	//compare arrays - only arrays of the same type can be compared
	var x = [...]int{1, 2, 3, 4}
	var y = [4]int{1, 2, 3, 4}
	fmt.Println(x == y)

}

func print(text string, x interface{}) {
	fmt.Println(text, " -> ", x)
}
