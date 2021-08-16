package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	simple()
	simple("dog", "cat", "horse")
	complex(1, "blue", true, 10.5, []string{"Chile", "Costa Rica"}, map[string]int{"id1": 100, "id2": 200})
	fmt.Println(it_is_slice("dog", "cat", "horse"))
}

func simple(s ...string) {
	fmt.Println(s)
}

func complex(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "is of type", reflect.ValueOf(v).Kind())
	}
}

func it_is_slice(params ...string) string {

	fmt.Printf(`
        slice    : %#v
        values   : %s
        length   : %d
        capacity : %d`+
		"\n\n",
		params,
		strings.Join(params, ", "),
		len(params),
		cap(params))

	return strings.Join(params, " ")

}
