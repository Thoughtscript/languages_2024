package main

import (
	g "review/generics"
	i "review/interfaces"
	o "review/ood"
	p "review/pointers"
)

func main() {
	i.InterfaceTest()

	g.EmptyInterfaceTest()
	g.GenericsTest()

	o.StructInheritanceTest()
	o.TypeInheritanceTest()

	p.PointersTestOne()
	p.PointersTestTwo()

}
