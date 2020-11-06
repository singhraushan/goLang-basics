package main

import "fmt"

func main() {
	simpleForloop()
	whileForloop()
	simpleRangeForloop()
	stringRangeForloop()
	mapRangeForloop()
	ChannelForloop()
	//infiniteForloop()

	//break,goto & continue
	loopBreak()
	loopGoto()
	loopContinue()
}
func simpleForloop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func infiniteForloop() {
	for {
		fmt.Println("infinite loop")
	}
}

func whileForloop() {
	i := 0
	for i < 5 {
		fmt.Println("whileForloopL: ", i)
		i += 1
	}
}

func simpleRangeForloop() {
	marks := []int{53, 65, 13, 95, 76}
	for i, mark := range marks {
		fmt.Printf("index i: %d & marks : %d\n", i, mark)
	}
}

func stringRangeForloop() {
	for i, v := range "Raushan" {
		fmt.Println(i, v)
	}
}

func mapRangeForloop() {
	mmap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	for key, value := range mmap {
		fmt.Println("key is:", key, "and value is:", value)
	}
}

func ChannelForloop() {
	chanl := make(chan int)
	go func() {
		chanl <- 10
		chanl <- 100
		chanl <- 1000
		chanl <- 10000
		close(chanl) //else deadlock
	}()

	for i := range chanl {
		fmt.Println("i is:", i)
	}
}

func loopBreak() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
}

func loopGoto() {
	i := 0

Label1:
	for ; i < 5; i++ {
		if i == 3 {
			i += 2
			goto Label1
		}
		fmt.Println("goto Label1: ", i)
	}
}

func loopContinue() {
	i := 0
	for ; i < 5; i++ {
		if i == 2 {
			i += 2
			continue
		}
		fmt.Println("continue: ", i)
	}
}
