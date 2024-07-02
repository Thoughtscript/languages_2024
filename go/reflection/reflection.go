package reflection

import (
	"fmt"
	"reflect"
	g "review/generics"
)

// w.r.t to Type Inheritance and Generics

type S interface{}

type W interface {
	M()
}

var stype = reflect.TypeOf((*S)(nil)).Elem() // Otain the type of S

var wtype = reflect.TypeOf((*W)(nil)).Elem() // Otain the type of S

type AA struct {
	msg string
}

type BB AA

type CC BB

func (p AA) M() {
	fmt.Println(1 + 2)
}

// Each of BB and CC require the implementation method M() they do not automatically or implicitly inherit like super()
func (p BB) M() {
	fmt.Println(1 + 2)
}

func (p CC) M() {
	fmt.Println(1 + 2)
}

// Note that the following won't allow T to be any of AA, BB, CC

// func exampleFour[T *S](s T, t T) {
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

// func exampleFour[T T.Implements(stype)](s T, t T) {
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

// func exampleFour[T.Implements(stype)](s T, t T) {
// 	fmt.Println(s)
// 	fmt.Println(t)
// }

func exampleFive(s AA, t any) {
	T := reflect.TypeOf(t)
	if T.Implements(stype) {
		fmt.Println("reflection > exampleFive", s, t)
	} else {
		panic("reflection > t doesn't implement S")
	}
}

func exampleSix(s AA, t S) {
	fmt.Println("reflection > exampleSix", s, t)
}

func exampleSeven(s AA, t W) {
	fmt.Println("reflection > exampleSeven", s, t)
}

func exampleEight(s AA, t any) {
	T := reflect.TypeOf(t)
	if T.Implements(wtype) {
		fmt.Println("reflection > exampleEight", s, t)
	} else {
		panic("reflection > t doesn't implement W")
	}
}

// Note that the above is not the same as the following:
func exampleNine[TT W](s TT, t TT) {
	fmt.Println("reflection > exampleNine", s, t)
}

func ReflectionTest() {
	var a = AA{
		msg: "a",
	}

	var b = BB{
		msg: "b",
	}

	var c = CC{
		msg: "c",
	}

	// Note that inheritance like so isn't obeyed
	// and even in the case where both TypeB and TypeC share a common interface

	// exampleOne(s, u)
	// exampleOne(u, s)

	var A = reflect.TypeOf(a)
	var B = reflect.TypeOf(b)
	var C = reflect.TypeOf(c)
	fmt.Println("reflection >", A, B, C)

	fmt.Println("reflection >", stype)
	fmt.Println("reflection >", A.Implements(stype), B.Implements(stype), C.Implements(stype))

	exampleFive(a, c)
	exampleFive(a, "string will implement subtype!") // In Go 1.22 this will actually succeed as string implements stype!

	exampleSix(a, "string will implement subtype!") // In Go 1.22 this will actually succeed as string implements stype!
	exampleSix(a, b)

	x := [...]int{1, 2, 3, 4, 5}
	exampleFive(a, x)
	exampleSix(a, x)

	var d = g.TypeC{
		// msg: "d",
		// First this is strange that this can't
		// be picked up by the compiler and interpreter
	}

	exampleFive(a, d) // these succeed despite the attempt to restrict the above by Type
	exampleSix(a, d)  // these succeed despite the attempt to restrict the above by Type

	// So S is behaving as the Empty Interface even though it's named.
	// Observe that by using W instead (with an Implemented Method) we can restrict
	// These Types must implement M() to work
	// exampleSeven(a, d)
	exampleSeven(a, b)
	exampleSeven(a, c)

	// These aren't valid
	// exampleNine(a, c)
	// exampleNine(a, d)

	exampleNine(a, a) // this is though suggesting it only goes 1 implementation deep

	// These will now correctly Type check by implementing through W
	exampleEight(a, c)

	// De facto catch clause w/ recover()
	// https://go.dev/blog/defer-panic-and-recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic!", r)
		}
	}()
	exampleEight(a, d)
	// Careful with what goes underneath - should use the above in wrapping func or go routine!
}
