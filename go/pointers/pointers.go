package pointers

import (
	"fmt"
)

// https://go.dev/tour/moretypes/1
func PointersTestOne() {
	i, j := 42, 2701

	p := &i                       // point to i
	fmt.Println("pointers >", *p) // read i through the pointer
	*p = 21                       // set i through the pointer
	fmt.Println("pointers >", i)  // see the new value of i

	p = &j                       // point to j
	*p = *p / 37                 // divide j through the pointer
	fmt.Println("pointers >", j) // see the new value of j

	fmt.Println("pointers >", *p)
	fmt.Println("pointers >", &p)
}

func PointersTestTwo() {
	var num int = 100
	fmt.Println("pointers > <variable> 'num' has <value>:", num)

	var numAddress *int = &num
	fmt.Println("pointers > 'numAddress' obtains <pointer> via & <address operator> and <pointer type>: *int = &num ", numAddress)

	var derefNum int = *numAddress
	fmt.Println("pointers > <derefences> back to the <value> with <dereferencing operator> on 'numAddress': derefNum = *numAddress", derefNum)
	fmt.Println("pointers > call & <address operator> on 'derefNum' to obtain <address>: &derefNum", &derefNum)
	fmt.Println("pointers > can <dereference> back directly with & and * operators sequentially: *&derefNum", *&derefNum)

	*numAddress = 42
	fmt.Println("pointers > set <value> on <dereference> of 'numAddress', then get <address> from <pointer type>: *numAddress = 42", numAddress)
	fmt.Println("pointers > <dereference> back to <value>: *numAddress ", *numAddress)
}
