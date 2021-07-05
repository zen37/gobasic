package main

import "fmt"

func main() {

	//multidimensional arrays
	var am [2][4]int // this is an array of length 2, whose type is an array of ints of length 4
	fmt.Println("multimensional aray [2][4]int", am)
	var sm [][]int //simulate multidimensional slices and make a slice of slices:
	fmt.Println("multimensional slice [][]int", sm)

	var s1 [][]string

	x := [][]string{{"x", "y"}}
	s1 = append(s1, x...)
	fmt.Println("s1:", s1)

	s2 := [][]string{{"a", "b"}, {"c", "d"}, ""}
	fmt.Println("s2:", s2)
	fmt.Println("s2[0]:", s2[0])
	fmt.Println("s2[1][0]:", s2[1][0])

	fmt.Println("loop over:", s2)
	for k1, v1 := range s2 {
		fmt.Println("key=", k1, "value=", v1)
		for k2, v2 := range v1 {
			fmt.Println("key=", "[", k1, "]", "[", k2, "]", "value=", v2)
		}
	}

}
