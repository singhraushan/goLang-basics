package main

import (
	"fmt"
	"time"
)

func portal1(ch1 chan string) { //like list ArrayList<>  String
	time.Sleep(3 * time.Second)
	ch1 <- "Welcome to channel 1"
	fmt.Println("portal1 end")
}

func portal2(ch2 chan string) {
	time.Sleep(9 * time.Second)
	ch2 <- "Welcome to channel 2"
	fmt.Println("portal2 end")
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	//ch2 <- "afsgd" //you can't add value in channel without Goroutine(not main thread(Goroutine))
	go portal1(ch1)
	go portal2(ch2) //ch2 is like pointer/object so changes available at calling place

	select {
	case c1 := <-ch1:
		fmt.Println(c1)
	case c2 := <-ch2:
		fmt.Println(c2)
	}

	fmt.Println("========================")

	ch3 := make(chan string)
	ch4 := make(chan string)

	go portal3(ch3)
	go portal4(ch4)

	select { //it will print only one value, because it's not looping
	case c3 := <-ch3:
		fmt.Println(c3)
	case c4 := <-ch4:
		fmt.Println(c4)
	}

	fmt.Println("========================")

	selectWithDefaultCase()
	fmt.Println("========================")
	selectWithoutDefaultCase()
	fmt.Println("========================")
	selectWithoutCase() //No C.E. but deadloack because select waiting
}
func selectWithDefaultCase() {
	channel1 := make(chan int)
	select {
	case <-channel1:
	default:
		fmt.Println("==default=======selectWithDefaultCase====Not found=") // this get executed
	}
}
func selectWithoutCase() {
	select {}
}

//No C.E. but deadloack because select waiting
func selectWithoutDefaultCase() {
	mychannel := make(chan int)
	// channel is not ready
	// and no default case
	select {
	case <-mychannel:
	}
}

func portal3(ch3 chan string) {
	for i := 0; i < 5; i++ {
		ch3 <- "Welcome to channel 3-" + string(i)
	}
}

func portal4(ch4 chan string) {
	for i := 0; i < 5; i++ {
		ch4 <- "Welcome to channel 4-" + string(i)
	}
}
