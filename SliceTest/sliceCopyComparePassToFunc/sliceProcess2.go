package main

import "fmt"

func main() {
	copySlice()
	fmt.Println("-----------------------")
	copySlice2()
	fmt.Println("-----------------------")
	slice1 := []string{"Geeks", "for", "Geeks", "GFG"}
	fmt.Println("main: Before modifyScliceInFunc call slice1 :", slice1)
	modifyScliceInFunc(slice1)
	fmt.Println("main: After modifyScliceInFunc call slice1 :", slice1) // changes reflecting here. but in Array just opposite concept
	modifyScliceInFunc2(slice1)
	fmt.Println("main: After modifyScliceInFunc2 call slice1 :", slice1)
}

func modifyScliceInFunc(slice1 []string) {
	slice1[2] = "Java"
	fmt.Println("modifyScliceInFunc: After modification slice1 :", slice1)
}
func modifyScliceInFunc2(slice1 []string) {
	// Here we only modify the slice
	// Using append function
	// Here, this function only modifies
	// the copy of the slice present in
	// the function not the original slice
	slice1 = append(slice1, "Ruby") // this not on calling place
	fmt.Println("modifyScliceInFunc2: After modification slice1 :", slice1)
}

func copySlice() {
	var slice1 []int = []int{1, 2, 3, 4, 5, 6, 7}
	var slice2 []int
	slice3 := make([]int, 2)
	slice4 := []int{10, 20, 30, 40, 50}

	// Before copying
	fmt.Println("Slice_1:", slice1)
	fmt.Println("Slice_2:", slice2)
	fmt.Println("Slice_3:", slice3)
	fmt.Println("Slice_4:", slice4)

	fmt.Println("Total no. of copied element from Slice2(nil) to slice1:", copy(slice1, slice2))
	fmt.Println("After copying Slice1:", slice1)

	fmt.Println("Total no. of copied element from Slice1(7) to slice3(2):", copy(slice3, slice1))
	fmt.Println("After copying Slice3:", slice3)

	copy_3 := copy(slice4, slice1)
	fmt.Println("Total no. of copied element from Slice1(7) to slice4(5):", copy_3)
	fmt.Println("After copying Slice4:", slice4)

	copy_4 := copy(slice1, slice4)
	fmt.Println("Total no. of copied element from Slice4(5) to slice1(7):", copy_4)
	fmt.Println("After copying Slice1:", slice1)
}
func copySlice2() {
	slice1 := []string{"Geeks", "for", "Geeks", "GFG"}
	slice2 := make([]string, 3)

	// Before copying
	fmt.Println("Slice_1:", slice1)
	fmt.Println("Slice_2:", slice2)

	fmt.Println("Total no. of copied element from Slice1(4) to slice2(3):", copy(slice2, slice1))
	fmt.Println("After copying Slice2:", slice2)

	copy_2 := copy(slice1, []string{"123geeks", "gfg"})
	fmt.Println("Total no. of copied element :", copy_2)
	fmt.Println("After copying Slice1:", slice1)

}
