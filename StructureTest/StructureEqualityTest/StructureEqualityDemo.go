package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	name, city string
	pincode    int
}

type Author struct {
	name, city string
	pincode    int
}

var address1 Address
var address2 Address

func main() {
	a1 := Address{"Raushan", "Begusarai", 848204}
	a2 := Address{"Amit", "Begusarai", 848204}
	a3 := Address{"Raushan", "Begusarai", 848204}

	author := Author{"Raushan", "CSE", 4646}

	fmt.Println("a1 and a2 are equal: ", reflect.DeepEqual(a1, a2))
	fmt.Println("a1 and a2 are equal: ", a1 == a2)

	fmt.Println("a1 and a3 are equal: ", reflect.DeepEqual(a1, a3))
	fmt.Println("a1 and a3 are equal: ", a1 == a3)

	fmt.Println("a1 and author are equal: ", reflect.DeepEqual(a1, author)) //false
	//fmt.Println("a1 and author are equal: ", a1 == author)                  // mismatch type

	//only declared
	fmt.Println("address1 and address2 are equal: ", reflect.DeepEqual(address1, address2)) //true
	fmt.Println("address1 and address2 are equal: ", address1 == address2)                  // true
}
