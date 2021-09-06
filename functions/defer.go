package main

import "fmt"

func cleanup() error {
	fmt.Println("Running cleanup...")
	return fmt.Errorf("error on cleanup")
}

func getMessage() (string, error) {

	err := cleanup()

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
