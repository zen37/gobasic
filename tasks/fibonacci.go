package main

//https://blog.devgenius.io/big-int-in-go-handling-large-numbers-is-easy-157cb272dd4f

import (
	"fmt"
)

// returns a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		result := x
		x, y = y, x+y
		return result
	}
}

/* func fibonacci_big() func() big.Int {
	x := *big.NewInt(0)
	y := *big.NewInt(1)
	return func() big.Int {
		result := x
		x, y = y, x+y
		return result
	}
} */

func main() {
	f := fibonacci()
	for i := 0; i < 111; i++ {
		fmt.Println("i=", i, " ", f())
	}
}
