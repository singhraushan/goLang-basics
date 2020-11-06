package main

import "fmt"

func main() {
	fmt.Println("main start")
	fmt.Println("sum s: ", sum(2, 3))
	fmt.Println("main middle")
	defer fmt.Println("sum s: ", sum(10, 5)) // this excecute after all execution over
	fmt.Println("main end")

	fmt.Println("--------calling multiDeferTest-------")
	multiDeferTest()

}

func multiDeferTest() {
	defer fmt.Println("multiDeferTest start") //this excecute 3rd among all defer within multiDeferTest() using LIFO concept
	fmt.Println("sum 1 & 1: ", sum(1, 1))
	fmt.Println("multiDeferTest middle")
	defer fmt.Println("sum 10 & 1: ", sum(10, 1)) //this excecute 2nd among all defer within multiDeferTest() using LIFO concept
	fmt.Println("multiDeferTest end")
	defer fmt.Println("sum s: ", sum(100, 25)) //this excecute 1st among all defer within multiDeferTest() using LIFO concept
}

func sum(a, b int) int {
	return a + b
}
