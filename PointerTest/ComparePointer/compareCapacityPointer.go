package main

import "fmt"

func main() {
	compareTwoPointer()
	fmt.Println("================")
	pointerCapacity()
	fmt.Println("================")
	pointerLength()
}
func compareTwoPointer() {
	val1 := 2345
	val2 := 567
	p1 := &val1
	p2 := &val2
	p3 := &val1

	res1 := &p1 == &p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res1)

	fmt.Println("Is p1 pointer is equal to p2 pointer: ", p1 == p2)
	fmt.Println("Is p1 pointer is equal to p3 pointer: ", p1 == p3)
	fmt.Println("Is p2 pointer is equal to p3 pointer: ", p2 == p3)
	fmt.Println("Is p3 pointer is equal to p1 pointer: ", &p3 == &p1)

	fmt.Println("Is p1 pointer not equal to p2 pointer: ", &p1 != &p2)
	fmt.Println("Is p1 pointer not equal to p2 pointer: ", p1 != p2)
	fmt.Println("Is p1 pointer not equal to p3 pointer: ", p1 != p3)
	fmt.Println("Is p2 pointer not equal to p3 pointer: ", p2 != p3)
	fmt.Println("Is p3 pointer not equal to p1 pointer: ", &p3 != &p1)
}

func pointerCapacity() {
	var pt1 [5]*int
	var pt2 [7]*int
	var pt3 [9]*int
	fmt.Println("capacity of pt1: ", cap(pt1)) //5
	fmt.Println("capacity of pt2: ", cap(pt2)) //7
	fmt.Println("capacity of pt3: ", cap(pt3)) //9

	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	for x := 0; x < len(pt1); x++ {
		pt1[x] = &arr[x]
	}

	for x := 0; x < len(pt1); x++ {
		fmt.Printf("Value of pt1[%d] = %d\n", x, *pt1[x]) //9
	}
	fmt.Println("Capacity of arr: ", cap(arr))
	fmt.Println("Capacity of pt1: ", cap(pt1))
}

func pointerLength() {
	var pt1 [4]*int
	var pt2 [7]*int
	var pt3 [9]*int
	fmt.Println("length of pt1: ", len(pt1)) //4
	fmt.Println("length of pt2: ", len(pt2)) //7
	fmt.Println("length of pt3: ", len(pt3)) //9

	arr := [6]int{1, 2, 3, 4, 5, 6}

	for x := 0; x < len(pt1); x++ {
		pt1[x] = &arr[x]
	}

	for x := 0; x < len(pt1); x++ {
		fmt.Printf("Value of pt1[%d] = %d\n", x, *pt1[x]) //9
	}
	fmt.Println("length of arr: ", len(arr))
	fmt.Println("length of pt1: ", len(pt1))
}
