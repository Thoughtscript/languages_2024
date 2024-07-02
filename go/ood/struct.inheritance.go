package ood

import (
	"fmt"
)

type Animal struct {
	name string
}

// Weak Inheritance by Composition
type Mammal struct {
	Animal
}

// As a field
type Reptile struct {
	animalParts Animal
}

func StructInheritanceTest() {
	dog := Mammal{
		Animal{
			name: "fido",
		},
	}

	fmt.Println(dog)

	lizard := Reptile{
		animalParts: Animal{
			name: "lizzy",
		},
	}

	fmt.Println(lizard)
}
