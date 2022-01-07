package main

import (
	"fmt"
	"reflect"
)

type Request struct {
	Method   string
	Resource string //path
	Protocol string
}

type s struct {
	ID        int
	Title     string
	Request   Request
	Price     float64
	Interface interface{}
	Exists    bool
	Many      []string
}

func main() {
	s := s{}

	iterateStruct(s)

}

func iterateStruct(s interface{}) {

	e := reflect.ValueOf(s)

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varKind := e.Field(i).Kind()
		//fmt.Println(e.Type().Field(i).Name)
		if varKind == reflect.Struct {
			fmt.Println("-----------------------------------")
			fmt.Println("struct in struct: ", e.Type().Field(i).Name)
			iterateStruct(e.Field(i).Interface())
			fmt.Println("-----------------------------------")
		}
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v %v\n", varName, varKind, varType, varValue)
	}

}
