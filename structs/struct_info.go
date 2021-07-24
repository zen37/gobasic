package main

import (
	"fmt"
	"reflect"
)

type s struct {
	ID     int
	Title  string
	Price  float64
	Exists bool
	Many   []string
}

func main() {
	s := s{}
	e := reflect.ValueOf(s)

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}
