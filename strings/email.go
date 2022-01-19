package main

import (
	"fmt"
	"net/mail"
)

func main() {
	for _, email := range []string{
		"good@exmaple.com",
		"bad-example",
	} {
		fmt.Printf("%18s valid: %t\n", email, valid(email))
	}

}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
