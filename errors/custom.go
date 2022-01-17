package main

import (
	"errors"
	"fmt"
	"log"
)

//custom error type
type DivisionError struct {
	dividend int
	divisor  int
	msg      string
}

// DivisionError implements the error interface
func (e *DivisionError) Error() string {
	return e.msg
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{msg: fmt.Sprintf("cannot divide '%d' by zero", a)}
	}
	return a / b, nil
}

func main() {

	a, b := 10, 0
	result, err := divide(a, b)

	if err != nil {
		var divErr *DivisionError
		switch {
		//check and convert from standard error to our more specific DivisionError.
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is not mathematically valid: %s\n",
						divErr.dividend, divErr.divisor, divErr.msg)
			log.Fatalln("error is: ", err)

		default:
			fmt.Printf("unexpected divison error: %s\n", err)
			log.Fatalln(err)
		}
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)

}
