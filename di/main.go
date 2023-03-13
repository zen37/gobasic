package main

import (
	"log"
)

type Database interface {
	GetUser() string
	GetAllUsers() []string
}

type SQLDb struct{}

func (db SQLDb) GetUser() string {
	return "Joe"
}

type SQLDbBetter struct{}

func (db SQLDbBetter) GetUser() string {
	return "better sql db"
}

func (db SQLDbBetter) GetAllUsers() []string {
	return []string{}
}

type Greeter interface {
	Greet(userName string)
}

type NiceGreeter struct{}

func (greet NiceGreeter) Greet(userName string) {
	log.Printf("Hello %s! Nice to meet you", userName)
}

type Program struct {
	Db      Database
	Greeter Greeter
}

func (p Program) Execute() {
	user := p.Db.GetUser()
	p.Greeter.Greet(user)
}

func main() {
	//db := SQLDb{}
	db := SQLDbBetter{}
	greeter := NiceGreeter{}

	p := Program{db, greeter}

	p.Execute()
}
