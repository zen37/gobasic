package main

import (
	"fmt"
	"reflect"
)

func main() {
	simple()
	simple("dog", "cat", "horse")
	complex(1, "blue", true, 10.5, []string{"Chile", "Costa Rica"}, map[string]int{"id1": 100, "id2": 200})
}

func simple(s ...string) {
	fmt.Println(s)
}

func complex(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "is of type", reflect.ValueOf(v).Kind())
	}
}
