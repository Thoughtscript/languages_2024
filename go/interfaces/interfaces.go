package interfaces

import (
	"fmt"
)

// Uses implicit aggregation like so: ../../../java/src/main/java/thoughtscript/io/review/InterfaceAggregationTest.java#L7-L11
// Added GoLang 1.18
type Shape interface {
	AddCoordinates() float32 // Requires a Receiver Function to be defined by the relevant implementations
	// The above if defined like so cannot must not have a Pointer argument - .e.g (p *Point) and must have matching return type
}

// Any type below will be required to implement

// https://github.com/Thoughtscript/go_refresh/blob/master/examples/5-receivertype/receiver/Point.go
// Note this is a weakly OOP language
// Note the lack of (true, non-compositional) inheritance and constructors
type Point struct {
	X float32
	Y float32
	Z float32
}

// Pseudo constructor/factory
// Note capitalized name indicates "to export"
func NewPoint(X float32, Y float32, Z float32) *Point {

	// Partially initialize object
	p := Point{X: X, Y: Y}
	p.Z = Z

	// Return pointer
	return &p
}

// Pseudo-class methods
// Call by value
// Copy of value made

// This won't satisfy the Interface constraint
func AddCoordinates(p Point) float32 {
	return p.X + p.Y + p.Z
}

// Receiver type function
// This is also how visibility is controlled:
// By exporting receiver methods but limiting exporting of structs (by using lower-case names)

func (p Point) AdditionReceiver() float32 {
	return p.X + p.Y + p.Z
}

// The Interface Function implementation must be a Receiver
func (p Point) AddCoordinates() float32 {
	return p.X + p.Y + p.Z
}

func InterfaceTest() {
	// type in declaration can be the Interface
	var s Shape = Point{1, 2, 3}
	fmt.Println("interfaces >", s)

	a := Point{1, 2, 3}
	fmt.Println("interfaces >", a.AddCoordinates())
	fmt.Println("interfaces >", a)

	b := NewPoint(3, 4, 5)
	fmt.Println("interfaces >", b)
	fmt.Println("interfaces >", &b)
	fmt.Println("interfaces >", *b)

	c := AddCoordinates(a)
	fmt.Println("interfaces >", c)

	d := a.AdditionReceiver()
	fmt.Println("interfaces >", d)
}
