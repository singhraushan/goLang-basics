package main

import "fmt"

func main() {
	normalVariableNotHoldingAddressAndNotPointingAddress()
	fmt.Println("====================")
	pointerDeclarationAndInitialization()
	x := 15 // passing address of local variable to another method
	modifyValueAtPointer(&x)
	fmt.Println("value of x after ethod call==========", x) // value got changed
	fmt.Println("====================")
	pointerToFunction()
	fmt.Println("====================")
	n := rpf() //returning pointer
	fmt.Println("value at the address holded by n: ", *n)
}
func normalVariableNotHoldingAddressAndNotPointingAddress() {
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Println("Value of x ", x) //225
	fmt.Printf("Value of x in hexdecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x) //225

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Println("Value of y ", y) //156
	fmt.Printf("Value of y in hexdecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y) //156
}

func pointerDeclarationAndInitialization() {
	var x int = 45
	var p *int
	p = &x

	fmt.Println("x:", x)
	fmt.Println("adddress of x ", &x)
	fmt.Println("adddress of p ", &p)
	fmt.Println("adddress hoding p ", p)
	fmt.Println("value at address holded by p:", *p)

	var i *int
	var s *string
	var f *float64
	fmt.Println("default value of pointer i:", i) //nil
	fmt.Println("default value of pointer s:", s) //nil
	fmt.Println("default value of pointer f:", f) //nil

	//s=&x //C.E.(string pointer can't hold address of int variable. Pointer & variable data type should & must be same)
	y := 15 // or var y = 15
	var p2 = &y
	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p2 = ", p2)
}

func modifyValueAtPointer(i *int) {
	fmt.Println("Before modification: Value pointed by i ", *i)
	*i = 500
	fmt.Println("After modification: Value pointed by i ", *i)
}

func pointerToFunction() {
	var x = 3
	fmt.Println("value of x: ", x)
	ptf(&x) //1st way to call
	fmt.Println("After func call: value of x: ", x)
	var pt1 = &x
	*pt1 = 18
	fmt.Println("After func call1: value of x: ", x)
	ptf(pt1) // 2nd way to call
	fmt.Println("After func call2: value of x: ", *pt1)
}
func ptf(i *int) {
	*i = 200
}

func rpf() *int {
	x := 7
	return &x
}
