package main

import (
	"fmt"
	"time"
)

func main() {
	//C.E. because no sum is not in use. To avoid this use _(blank identifier)
	//sum, sub := cal(10, 5)

	_, sub := cal(10, 5)
	fmt.Print(sub, "\n")
	fmt.Println("current time: ", time.Now())
}

func cal(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}
