package main

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("main")
	a := 10
	fmt.Print("main: value of before modify:", a, "\n")
	modify(&a)
	fmt.Print("main: value of after modify: ", a)
}
func modify(x *int) {
	*x = 70
}

func init() {
	fmt.Println("init3")
}
