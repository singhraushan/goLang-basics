package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	arrayToSlice()
	fmt.Println("-------------------------")
	sliceUsingLiteral()
	fmt.Println("-------------------------")
	sliceToSlice()
	fmt.Println("-------------------------")
	sliceUingMakeFunc()
	fmt.Println("-------------------------")
	iterate()
	fmt.Println("-------------------------")
	importtantPointWithSlice()
}

func arrayToSlice() {
	arr1 := [5]string{"Go", "Java", "C", "Ruby", "Python"}
	slice1 := arr1[0:3]
	fmt.Println("slice1: ", slice1)
	fmt.Println("length of slice1: ", len(slice1))   //3
	fmt.Println("capacity of slice1: ", cap(slice1)) //5
	slice2 := arr1[2:3]
	slice3 := arr1[0:] //default(len(arr1) as high bound)
	slice4 := arr1[:4] // default(0 as low bound)
	slice5 := arr1[:]  // default(0 as low bound & len(arr1) as high bound) all element copied
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice5: ", slice5)

}

func sliceUsingLiteral() {
	slice1 := []string{"Go", "Java", "C", "Ruby", "Python"}
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("My Slice1:", slice1)
	fmt.Println("My Slice2:", slice2)
}

func sliceToSlice() {
	original_slice := []int{1, 2, 3, 4, 5}
	new_slice1 := original_slice[2:3]
	new_slice2 := original_slice[0:] //default(len(arr1) as high bound)
	new_slice3 := original_slice[:4] // default(0 as low bound)
	new_slice4 := original_slice[:]  // default(0 as low bound & len(arr1) as high bound) all element copied
	fmt.Println("New slice 1: ", new_slice1)
	fmt.Println("New slice 2: ", new_slice2)
	fmt.Println("New slice 3: ", new_slice3)
	fmt.Println("New slice 4: ", new_slice4)
}
func sliceUingMakeFunc() {
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	slice3 := make([]int, 0)

	fmt.Println("New slice 1 using make(): ", slice1, "capacity:", cap(slice1))
	fmt.Println("New slice 2 using make(): ", slice2, "capacity:", cap(slice2))
	fmt.Println("New slice 3 using make(): ", slice3, "capacity:", cap(slice3))
}

func iterate() {
	myslice := []string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	for i := 0; i < len(myslice); i++ {
		fmt.Print(myslice[i], " ")
	}
	fmt.Println()
	for index, val := range myslice {
		fmt.Print(index, " ", val, " ")
	}
	fmt.Println()
	for _, val := range myslice {
		fmt.Print(val, " ")
	}
}
func importtantPointWithSlice() {
	var sl1 []int
	fmt.Println("sl1: ", sl1)
	fmt.Println("len(sl1): ", len(sl1))
	fmt.Println("cap(sl1): ", cap(sl1))

	arr1 := [5]string{"Go", "Java", "C", "Ruby", "Python"}
	sl2 := arr1[0:3]

	fmt.Println("Before:Original_Array: ", arr1)
	fmt.Println("Before:Original_Slice: ", sl2)
	sl2[0] = "C#"
	sl2[1] = "Swift"
	sl2[2] = "R"
	fmt.Println("After:Original_Array: ", arr1) // array also changing
	fmt.Println("After:Original_Slice: ", sl2)

	//Comparison of Slice
	sl3 := arr1[0:2]
	//fmt.Println(sl2 == sl3)//C.E. can't compare two sclice using == operator
	fmt.Println("sl2 and sl3 is equal: ", reflect.DeepEqual(sl2, sl3))
	fmt.Println("sl1 is nil: ", sl1 == nil)
	fmt.Println("sl3 is nil: ", sl3 == nil)

	//Multi Dimentional slice
	sl4 := [][]int{{1, 2}, {3, 4}}
	fmt.Println("sl4: ", sl4)

	//Sorting of Slice
	sl5 := []string{"Python", "Java", "C#", "Go", "Ruby"}
	sl6 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
	sl7 := []float64{45.47, 6.7, 23, 90, 33, 2.1, 56, 78.23, 89}

	fmt.Println("Before sorting:")
	fmt.Println("Slice 5: ", sl5)
	fmt.Println("Slice 6: ", sl6)
	fmt.Println("Slice 7: ", sl7)

	// Performing sort operation on the
	// slice using sort function
	sort.Strings(sl5)
	sort.Ints(sl6)
	sort.Float64s(sl7)

	fmt.Println("\nAfter sorting:")
	fmt.Println("Slice 5: ", sl5)
	fmt.Println("Slice 6: ", sl6)
	fmt.Println("Slice 7: ", sl7)
}
