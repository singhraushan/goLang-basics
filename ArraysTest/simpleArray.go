package main

import (
	"fmt"
	"reflect"
)

func main() {

	//declaration using var
	var arr1 [3]string
	arr1[0] = "Raushan"
	arr1[1] = "Terence"
	arr1[2] = "Janmejay"
	fmt.Println(arr1[0], arr1[1], arr1[2])

	//short hand declaration
	arr2 := [3]string{"Raushan", "Terence", "Janmejay"}
	fmt.Println(arr2[0], arr2[1], arr2[2])

	//2d Array:
	arr3 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
	}
	fmt.Println()
	//2d Array:
	var arr4 [3][2]int
	k := 0
	for i := 0; i < len(arr4); i++ { //++i(pre increment) not working here showing C.E.
		for j := 0; j < len(arr4[i]); j++ {
			arr4[i][j] = k //k++ && ++k not working here showing C.E.
			k++
		}
	}

	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			fmt.Print(arr4[i][j], " ")
		}
	}
	fmt.Println()
	var myarr [2]int
	fmt.Println("Elements of the Array :", myarr)

	//short hand declraration of array with 3 dot(...) inside []. if ellipsis ‘‘…’’(called ellipsis)
	var myarr2 [3]int = [...]int{1, 2, 3}
	fmt.Println("Elements of the Array :", myarr2)

	//length of array is determined based on value passed to array using ellipsis(...)
	myarr3 := [...]string{"Amit", "Shubham", "Suman"}
	fmt.Println("Elements of the Array :", myarr3)
	fmt.Println("length of the Array3 :", len(myarr3))

	//an array is of value type not of reference type(i.e. object kind)
	myarr4 := [5]int{1, 2, 3, 4, 5}
	new_myarr := myarr4
	myarr4[0] = 10     //this won't change new_myarr
	new_myarr[4] = 500 //this won't change myarr4
	fmt.Println("Elements of the myarr4 :", myarr4)
	fmt.Println("length of the new_myarr :", new_myarr)

	//compare two arrays using == operator
	Arr1 := [3]int{9, 7, 6}
	Arr2 := [...]int{9, 7, 6}
	Arr3 := [3]int{9, 1, 2}
	Arr4 := [3]string{"9", "7", "6"}
	if Arr1 == Arr2 {
		fmt.Println("Arr1 and Arr2 are equal")
		fmt.Println("reflect.DeepEqual(Arr1, Arr2)", reflect.DeepEqual(Arr1, Arr2))
	} else {
		fmt.Println("Arr1 and Arr2 are NOT equal")
	}

	if Arr1 == Arr3 {
		fmt.Println("Arr1 and Arr3 are equal")
	} else {
		fmt.Println("Arr1 and Arr3 are NOT equal")
	}

	if reflect.DeepEqual(Arr1, Arr4) { // Arr1 == Arr4 gives C.E.
		fmt.Println("Arr1 and Arr4 are equal")
	} else {
		fmt.Println("Arr1 and Arr4 are NOT equal")
	}

	fmt.Println("----------array copy by refe: -----")
	//array copy
	copyByRef()

	fmt.Println("----------pass array to func: -----")
	a1 := []int{1, 2, 3, 4, 5, 6}                                 // means this is slice NOT array
	fmt.Println("acceptUnsizedArray(a1)", acceptUnsizedArray(a1)) // pasing slice to func who accept slice
	//fmt.Println("acceptUnsizedArray(a1)", acceptSizedArray(a1))// C.E. can't pass slice for array
	a2 := [...]int{1, 2, 3, 4, 5} // this is array not slice
	//fmt.Println("acceptUnsizedArray(a1)", acceptUnsizedArray(a2)) //accept only slice, so can't pass array.
	fmt.Println("a2 before acceptUnsizedArray(a2)", a2)
	fmt.Println("acceptUnsizedArray(a2)", acceptSizedArray(a2))
	fmt.Println("a2 after acceptUnsizedArray(a2)", a2)
}

func copyByRef() {
	arr1 := [3]int{9, 7, 6}
	new_arr := &arr1

	fmt.Println("Before:Elements of the arr1 :", arr1)
	fmt.Println("Before:Elements of the new_arr :", *new_arr)
	arr1[0] = 600    // chnages reflect in new_arr also
	new_arr[2] = 900 // chnages reflect in arr1 also

	fmt.Println("After:Elements of the arr1 :", arr1)
	fmt.Println("After:Elements of the new_arr :", *new_arr)
}

func acceptUnsizedArray(arr1 []int) int { // means accept slice
	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	return sum
}

func acceptSizedArray(arr1 [5]int) int { // means accept array not slice
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr1[i]
	}
	arr1[0] = 500 // this not reflecting at calling place. this is only local change
	return sum
}
