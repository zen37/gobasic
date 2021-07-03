package main

import "fmt"

func main() {
	define()
	declare_slice()
	make_slice()
}

func print(text string, x interface{}) {
	fmt.Println(text, " -> ", x)
}

func declare_slice() {
	//create slice
	//using an empty slice literal:
	// 		this creates a zero-length slice, which is non-nil (comparing it to nil returns false).

	var y = []int{}
	printSlice(y)
	fmt.Println(y == nil) // false

	// without declaring a literal
	var x []int
	printSlice(x)
	fmt.Println(x == nil) //true

	x = append(x, 10)
	printSlice(x)

	x = append(x, 20)
	printSlice(x)

	x = append(x, 30)
	printSlice(x)

	x = append(x, 40)
	printSlice(x)

	x = append(x, 50)
	printSlice(x)

	//	Modifying the elements of a slice will also modify its underlying array:
	nums := [6]int{1, 2, 3, 4, 5}
	print("nums:= [6]int{1, 2, 3, 4, 5} is: ", nums)
	s := nums[0:4]
	print("nums[0:4] is: ", s)
	// modify s
	s[0] = 999
	print("after modifying s[0] nums[0:4] is: ", s)
	// which also modifies nums
	// => [999 2 3 4 5 6]
	print("nums:= [6]int{1, 2, 3, 4, 5} is: ", nums)

}

func printSlice(x []int) {
	fmt.Println("slice is: ", x, "len =", len(x), "cap =", cap(x))
}

func define() {
	// the array type is defined by its size + type of elements
	// by default an array is zero valued for ints is  0
	var a1 [3]int
	print("array [3]int", a1)
	var s1 []int
	print("slice []int", s1)
	print("[]int == nil", s1 == nil)

	var a2 = [3]int{1, 2, 3}
	fmt.Println("array [3]int{1, 2, 3} -> ", a2)
	var s2 = []int{1, 2, 3}
	fmt.Println("slice []int{1, 2, 3} -> ", s2)

	//sparse array, most elements are set to their 0 values, can specify only the indices with values
	var a3 = [7]int{1, 4: 7, 8}
	print("array [7]int{1, 4: 7, 8}", a3)
	var s3 = []int{1, 4: 7, 8}
	print("slice []int{1, 4: 7, 8}", s3)

	//compare arrays - only arrays of the same type can be compared
	var x = [...]int{1, 2, 3, 4}
	var y = [4]int{1, 2, 3, 4}
	fmt.Println(x == y)
	//compare slices - the onluy thing you can compare a slice with is nil
	print("s1 is []int, s1==nil", s1 == nil)
	print("s2 is []int{1, 2, 3}, s2==nil", s2 == nil)
}

/*
https://learning.oreilly.com/library/view/learning-go/9781492077206/ch03.html
If you have a good idea of how large your slice needs to be,
but don’t know what those values will be when you are writing the program, use make.
*/
// make, allows to specify the type, length, and, optionally, the capacity.
func make_slice() {

	fmt.Println("Use a non-zero length or zero length and capacity?")
	fmt.Println(" - If you are using a slice as a buffer, then specify a nonzero length.")

	opt2 := ` - If you are sure you know the exact size you want,
	you can specify the length and index into the slice to set the values.
	This is often done when transforming values in one slice and storing them in a second.
	The downside to this approach is that if you have the size wrong,
	you’ll end up with either zero values at the end of the slice or
	a panic from trying to access elements that don’t exist.`

	fmt.Println(opt2)

	opt3 := ` - In other situations, use make with a zero length and a specified capacity. 
	This allows you to use append to add items to the slice. 
	If the number of items turns out to be smaller, you won’t have an extraneous zero value at the end. 
	If the number of items is larger, your code will not panic.`

	fmt.Println(opt3)

	s := `Conclusion: The Go community is split between the second and third approaches.
	I [Jon Bodner] personally prefer using append with a slice initialized to a zero length.
	It might be slower in some situations, but it is less likely to introduce a bug.`

	fmt.Println(s)
	fmt.Println()

	//make an int slice of 3 elements that are 0, len is 3 and capacity is 3
	x := make([]int, 3)
	print2(x, "make([]int, 3)")

	// 10 is appended to the end of the slice, len and cap are increased
	x = append(x, 10)
	print2(x, "append(x, 10)")

	// specify initial capacity
	x = make([]int, 3, 32)
	print2(x, "make([]int, 3, 32)")

	// length 0
	x = make([]int, 0, 10)
	print2(x, "make([]int, 0, 10)")

	x = append(x, 1, 2, 3)
	print2(x, "append(x, 1, 2, 3)")

}

func print2(x []int, descr string) {
	fmt.Println(descr+" - "+"slice is: ", x, "len =", len(x), "cap =", cap(x))
}
