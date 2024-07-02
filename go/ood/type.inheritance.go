package ood

import "fmt"

// Cannot be basic type implementations and be used in type parameterization
// But can be used in custom types otherwise
type CustomString string

type B CustomString

func TypeInheritanceTest() {
	var AA CustomString = "AA"
	var BB B = "BB"
	fmt.Println("type inheritance >", AA)
	fmt.Println("type inheritance >", BB)
}
