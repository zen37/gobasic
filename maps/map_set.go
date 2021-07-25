package main

/*
Book Learning Go by Jon Bodner
Go doesn’t include a set, but you can use a map to simulate some of its features.
Use the key of the map for the type that you want to put into the set and use a bool for the value.


Some people prefer to use struct{} for the value when a map is being used to implement a set.
The advantage is that an empty struct uses zero bytes, while a boolean uses one byte.
The disadvantage is that using a struct{} makes your code more clumsy.
You have a less obvious assignment, and you need to use the comma ok idiom to check if a value is in the set:

Unless you have very large sets, it is unlikely that the difference in memory usage is significant enough to outweigh the disadvantages.
*/

import "fmt"

var vals []int

func init() {
	vals = []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
}

func main() {
	set_bool()
	set_empty_struct()
}

func set_bool() {
	intSet := map[int]bool{}
	//intSet := make(map[int]bool)
	fmt.Println(intSet)

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[5] {
		fmt.Println("5 is in the set")
	}
}

func set_empty_struct() {
	//intSet := map[int]struct{}{}
	intSet := make(map[int]struct{})
	fmt.Println(intSet)

	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	fmt.Println(intSet)

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
