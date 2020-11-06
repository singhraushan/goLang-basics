package main

import "fmt"

func main() {
	ifTest()
	ifElseTest()
	ifElseIFTest()
}

func ifTest() {
	var v int = 700

	if v < 1000 {
		fmt.Println("v is less than 1000")
	}
	fmt.Printf(" value of v is: %d\n", v)
}

func ifElseTest() {
	var v int = 1200

	if v < 1000 {
		fmt.Println("v is less than 1000")
	} else {
		fmt.Printf("v is greater than 1000\n")
	}
}

func ifElseIFTest() {
	var v int = 700

	if v == 200 {
		fmt.Println("v is equals to 200")
	} else if v == 100 {
		fmt.Println("v is equals to 100")
	} else if v == 400 {
		fmt.Println("v is equals to 400")
	} else {
		fmt.Println("value is not matching anyone of above")
	}
}
