//https://trstringer.com/golang-deferred-function-error-handling/

package main

import "fmt"

func cleanup() error {
	_, err := fmt.Println("Running cleanup...")
	return err
	//return fmt.Errorf("error from cleanup")
}

func doThis() error {
	return fmt.Errorf("error from doing another thing")
}

func getMessage() (msg string, errs []error) {

	defer func() {
		if err := cleanup(); err != nil {
			errs = append(errs, err)
		}

	}()
	if err := doThis(); err != nil {
		errs = append(errs, err)
	}

	return "hello world", errs
}

func main() {
	/*
		message, err := getMessage()
		if err != nil {
			fmt.Printf("Error getting message: %v\n", err)
		} else {
			fmt.Printf("Success. Message: '%s'\n", message)
		}
	*/
	message, errs := getMessage()
	if errs != nil {
		fmt.Printf("There are %d error(s)\n", len(errs))
		for _, err := range errs {
			fmt.Printf("Error: %v\n", err)
		}
	} else {
		fmt.Printf("Success. Message: '%s'\n", message)
	}
}
