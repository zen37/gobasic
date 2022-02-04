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
	s := s{ID: 10, Price: 32.0}
	//s := s{}

	iterateStruct(s)

}

func iterateStruct(s interface{}) {

	e := reflect.ValueOf(s)
	fmt.Println("reflect.ValueOf(<input>)", e)
	fmt.Println("<input> kind is:", e.Kind())

	/*
		In case we pass a pointer to reflect.ValueOf, and then attempt to call NumField(), we will end up with
		panic: reflect: call of reflect.Value.NumField on ptr Value
		In order to resolve this, we can check the message_value.Kind() for a pointer result,
		and then get the actual value, resolving the pointer:
	*/

	if e.Kind() == reflect.Ptr {
		e = e.Elem()
	}

	/*
		Calling e.NumField() will now report the correct number of fields within this struct.
		We can use this value to iterate over each one and get the needed names and values that we require.
	*/

	for i := 0; i < e.NumField(); i++ {
		fmt.Println("i=", i)
		varName := e.Type().Field(i).Name
		fmt.Println("name=", varName)
		varKind := e.Field(i).Kind()
		fmt.Println("kind=", varKind)
		if varKind == reflect.Struct {
			fmt.Println("-----------------------------------")
			fmt.Println("struct in struct: ", e.Type().Field(i).Name)
			iterateStruct(e.Field(i).Interface())
			fmt.Println("-----------------------------------")
		}
		varType := e.Type().Field(i).Type
		fmt.Println("type=", varName)
		varValue := e.Field(i).Interface()
		fmt.Println("value=", varValue)
		fmt.Printf("%v %v %v %v\n", varName, varKind, varType, varValue)
	}

}
