package pointers

import (
	"fmt"
)

// https://go.dev/tour/moretypes/1
func PointersTestOne() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(*p)
	fmt.Println(&p)
}

func PointersTestTwo() {
	var num int = 100
	fmt.Println("<variable> 'num' has <value>:", num)

	var numAddress *int = &num
	fmt.Println("'numAddress' obtains <pointer> via & <address operator> and <pointer type>: *int = &num ", numAddress)

	var derefNum int = *numAddress
	fmt.Println("<derefences> back to the <value> with <dereferencing operator> on 'numAddress': derefNum = *numAddress", derefNum)
	fmt.Println("call & <address operator> on 'derefNum' to obtain <address>: &derefNum", &derefNum)
	fmt.Println("can <dereference> back directly with & and * operators sequentially: *&derefNum", *&derefNum)

	*numAddress = 42
	fmt.Println("set <value> on <dereference> of 'numAddress', then get <address> from <pointer type>: *numAddress = 42", numAddress)
	fmt.Println("<dereference> back to <value>: *numAddress ", *numAddress)
}
