package main

import "fmt"

/*
value methods can be invoked on values and pointers

pointers methods can only be invoked on pointers as they can modify the receiver
invoking them on a copy of the value would cause those modifications to be discarded
*/

type struct_without_attr struct {
}

type struct_without_attr_2 struct {
}

type struct_with_attr struct {
	name  string
	value int
}

type struct_with_attr_2 struct {
	name  string
	value int
}

func (*struct_without_attr) String() string {
	return fmt.Sprintf("I am String() method, therefore the type implements Stringer()")
}

func (struct_without_attr_2) String() string {
	return fmt.Sprintf("I am String() method, therefore the type implements Stringer()")
}

func (s struct_with_attr) String() string {
	return fmt.Sprintf("attribute name is %s, attribute value is %d!", s.name, s.value)
}

func (s *struct_with_attr_2) String() string {
	s.name = "new name"
	return fmt.Sprintf("attribute name is %s, attribute value is %d!", s.name, s.value)
}

func main() {

	v1 := struct_without_attr{}
	fmt.Println("pointer receiver")
	fmt.Println("v1 := struct_without_attr{}, printing $v1:", &v1)
	fmt.Println("v1 := struct_without_attr{}, printing v1:", v1)

	fmt.Println()
	fmt.Println("value receiver")
	v2 := struct_without_attr_2{}
	fmt.Println("value receiver")
	fmt.Println("v2 := struct_without_attr_2{}, printing $v2:", &v2)
	fmt.Println("v2 := struct_without_attr_2{}, printing v2:", v2)

	fmt.Println()
	fmt.Println("pointer receiver")
	s := struct_with_attr{"name", 100}
	fmt.Println("s is: ", s)
	fmt.Println("&s is: ", &s)

	fmt.Println()
	fmt.Println("value receiver")
	s2 := struct_with_attr_2{"name2", 100}
	fmt.Println("s2.name before calling String():", s2.name)
	fmt.Println("s2 is: ", s2)
	fmt.Println("&s2 is: ", &s2)
	fmt.Println("s2.name after calling String():", s2.name)

}
