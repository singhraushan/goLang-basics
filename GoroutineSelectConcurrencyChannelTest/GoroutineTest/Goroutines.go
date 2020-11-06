package main

import (
	"fmt"
	"time"
)

func goTest(str string) {
	for x := 0; x < len(str); x++ {
		fmt.Println(str, x)
	}
}

func goTest2(str string) {
	for x := 0; x < 5; x++ {
		fmt.Println(str, x)
		time.Sleep(100)
	}
}

func main() {
	go goTest("Welcome")
	goTest("Raushan")
	fmt.Println("========================================")
	go goTest2("Welcome@@@@@@@")
	goTest2("Raushan@@@@@@@@@")

	fmt.Println("========================================")
	go func() {
		fmt.Println("Welcome!! to GeeksforGeeks")
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("========================================")
	multiGoroutine()
	fmt.Println("GoodBye!! to Main function")
}

func multiGoroutine() {
	fmt.Println("Welcome!! to multiGoroutine function")
	go aId()
	go aName()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("GoodBye!! to multiGoroutine function")
}

func aId() {
	ids := [4]int{1, 2, 3, 4}
	for x := 0; x <= 3; x++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("id: ", ids[x])
	}
}
func aName() {
	names := [4]string{"Rohit", "Suman", "Aman", "Ria"}
	for x := 0; x <= 3; x++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("name: ", names[x])
	}
}
