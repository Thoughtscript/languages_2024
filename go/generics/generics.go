package generics

import (
	"fmt"
)

// Cannot use Custom Types for Type Parameterization here
// that inherit from Basic Types
type TypeA struct {
	msg string
}

type TypeB TypeA

type TypeC TypeB

// with Type parameters
func exampleOne[T TypeA](s T, t T) {
	fmt.Println(s)
	fmt.Println(t)
}

func exampleTwo[T string](s T, t T) {
	fmt.Println(s)
	fmt.Println(t)
}

func anyParamTypeOne[T any](x T) T {
	return x
}

// any without Type parameters
func anyParamTypeTwo(x any) any {
	return x
}

// Accepts either type
func GenericsTest() {
	var s = TypeA{
		msg: "s",
	}

	var t = TypeA{
		msg: "t",
	}

	exampleOne(s, t)

	exampleTwo("I'm first", "I'm last")

	x := anyParamTypeOne("I'm x")
	fmt.Println(x)

	y := anyParamTypeTwo("I'm y")
	fmt.Println(y)
}
