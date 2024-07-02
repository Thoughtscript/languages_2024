package generics

import (
	"fmt"
)

type A struct{}

type B string

func golangGeneric(val interface{}) {
	fmt.Println("empty interface >", val)
}

// Accepts either type
func EmptyInterfaceTest() {
	var a A
	golangGeneric(a)

	var b B
	golangGeneric(b)
}
