//https://ibraheem.ca/writings/go-in-twenty/

package main

type Signed interface {
	isStrictlyNegative() bool
}

type Number struct {
	value int
}

//a struct implements/satisfies an interface if it has all of its methods
func (n *Number) isStrictlyNegative() bool {
	return n.value < 0
}

/*
Number doesn't implement Signed because isStrictlyNegative is defined only on *Number. This is because Number and *Number are different types:

// this will not compile
SignedIsStrictlyNegative(Number{})

// but this will
SignedIsStrictlyNegative(&Number{})
*/
func main() {
	SignedIsStrictlyNegative(&Number{})
}

func SignedIsStrictlyNegative(s Signed) bool {
	// call a method on the underlying type of `s`
	return s.isStrictlyNegative()
}
