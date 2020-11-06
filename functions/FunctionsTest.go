package main

import (
	"fmt"
	"strings"
)

func main() {
	p, q := 2, 3
	fmt.Printf("Before swap value of p & q are: %d and %d\n", p, q)
	callByValueSwap(p, q)
	fmt.Printf("After swap value of p & q are: %d and %d\n", p, q)
	fmt.Println("--------------------------------")
	callByRefSwap(&p, &q)
	fmt.Printf("After swap using pointer, value of p & q are: %d and %d\n", p, q)

	fmt.Println("--------------------------------")
	//variadic Functions
	//passing zero value
	fmt.Println(variadicFunctions())

	//passing multiple args
	fmt.Println(variadicFunctions("abc"))
	fmt.Println(variadicFunctions("abc", "xyz", "alpha"))

	//passing slice
	s := []string{"ab", "bc", "ca", "de"}
	fmt.Println(variadicFunctions(s...)) //pass with 3 dot(.) after slice name

	//anonymous funtion(without name also called as function literal in go)
	func() { // no name only () with func keyword
		fmt.Println("anonymous function")
	}() // calling using ()

	value := func() { // no name only () with func keyword
		fmt.Println("anonymous function with variable assignment")
	}
	fmt.Println("calling anonymous after here...........")
	value() // calling anonymous func using variable name with ()

	// Passing arguments in anonymous function
	func(input string) {
		fmt.Println(" Passing arguments in anonymous function ...........", input)
	}("Raushan") // calling from here

	// return of a function as anonymous function type
	anonymousVar := func(p, q string) string {
		return p + "-" + q
	}
	anonymousTypeArgFunctions(anonymousVar)

	anonymousVar2 := anonymousTypeReturnFunctions()
	fmt.Println("anonymousVar2: ", anonymousVar2("a", "b"))

	fmt.Println("---------------multi return---------")
	r1, r2 := multiReturnByTypeFunctions(2, 3)
	fmt.Println("return1 and return2: ", r1, r2)
	r3, r2 := multiReturnByTypeAndNameFunctions(2, 3)
	fmt.Println("return3 and return2: ", r3, r2)

	r4, r5 := namedReturnParametersFunctions(5, 10)
	fmt.Println("namedReturnParametersFunctions return4 and return5: ", r4, r5)

	r6, r7 := namedReturnParametersDefaultZerovalueFunctions(5, 10)
	fmt.Println("namedReturnParametersDefaultZerovalueFunctions return6 and return7: ", r6, r7) //print 0,0

	r8, r9 := namedReturnParametersDefaultValueFunctions("abc", "xyz")
	fmt.Println("namedReturnParametersDefaultValueFunctions return8 and return9: ", r8, r9, "abc") //print 0,0
}

func callByValueSwap(a, b int) int {
	c := a
	a = b
	b = c
	return c
}

func callByRefSwap(a, b *int) int {
	//a = &p i.e. a & p pointing to same location
	//b = &q i.e. b & q pointing to same location
	c := *a // c=2
	fmt.Println("c is holding value of a NOT ADDRESS: ", c)
	*a = *b // a=3
	*b = c  //b = 2
	return c
}

func variadicFunctions(input ...string) string {
	return strings.Join(input, "-")
}

// arg of a function as anonymous function type
func anonymousTypeArgFunctions(i func(a, b string) string) {
	fmt.Print("return of a function as anonymous function type: ", i("1st arg", "2nd arg")) //no line change
	fmt.Println()                                                                           //line change
}

// Returning anonymous function
func anonymousTypeReturnFunctions() func(a, b string) string {
	ret := func(a, b string) string {
		return "Returning anonymous function: a: " + a + " --b: " + b
	}
	return ret
}

func multiReturnByTypeFunctions(a, b int) (int, int) {
	return a + b, a - b
}

func multiReturnByTypeAndNameFunctions(a, b int) (x int, y int) {
	return a + b, a - b
}

func namedReturnParametersFunctions(a, b int) (x, y int) {
	//x := a + b // C.E.(Because x and y already initialize)
	x = a + b
	y = a - b
	return // here no need to mention x & y seperately
}

func namedReturnParametersDefaultZerovalueFunctions(a, b int) (x, y int) {
	//return deault x & y value as 0
	return // here no need to mention x & y seperately
}

func namedReturnParametersDefaultValueFunctions(a, b string) (x, y string) {
	//return deault x & y value as 0
	return // here no need to mention x & y seperately
}
