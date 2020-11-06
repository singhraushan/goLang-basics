package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	updateArray(&arr)
	fmt.Println("updated array:", arr) // reflecting at calling place
	slc := arr[:]
	updateSlice(slc)                   // pasing slice
	fmt.Println("updated array:", arr) // array also changed
	fmt.Println("updated slice:", slc) //reflecting at calling place
}

//not recomened to use this way. use slice. Because not good as readablity
func updateArray(arr *[5]int) { // NOT [5]*int. * should be in begining
	(*arr)[4] = 15 //NOT arr[4]. get whole array pointer than position i.e. (*arr)[4]
}

func updateSlice(arr []int) { // no pointer
	arr[4] = 25
}
