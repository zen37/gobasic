package main

import "fmt"

func cleanup() error {
	fmt.Println("Running cleanup...")
	return nil
}

func doThis() error {
	return fmt.Errorf("error on cleanup")
}

func getMessage() (msg string, err error) {

	defer func() {
		err = cleanup()
	}()
	err = doThis()

	return "hello world", err
}

func main() {
	message, err := getMessage()
	if err != nil {
		fmt.Printf("Error getting message: %v\n", err)
	} else {
		fmt.Printf("Success. Message: '%s'\n", message)
	}
}
