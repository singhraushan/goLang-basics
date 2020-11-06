package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	byteSclicesCompare()
	fmt.Println("================================")
	stringCompareUsingByteSlice()
	fmt.Println("================================")
	byteSclicesEqualityCheck()
	fmt.Println("================================")
	intSortOperation()
	fmt.Println("================================")
	stringSortOperation()
}
func byteSclicesCompare() {
	s1 := []byte{'A', 'B', '1', 'Z', '0'}
	s2 := []byte{'A', 'b', '1', 'Z', '0'}
	s3 := []byte{'A', 'B', '1', 'Z', '0'}

	var result1 int = bytes.Compare(s1, s2)
	if result1 == 0 {
		fmt.Println("s1 and s2 are equals")
	} else {
		fmt.Println("s1 and s2 are not equals")
	}
	if bytes.Compare(s1, s3) == 0 {
		fmt.Println("s1 and s3 are equals")
	} else {
		fmt.Println("s1 and s3 are not equals")
	}

	//using string you can create slice of byte. Then can use Compare(). Since Compare is not in strings
	s4 := []byte("Java") //Use () NOT {}
	s5 := []byte("JAVA") //Use () NOT {}
	if bytes.Compare(s4, s5) == 0 {
		fmt.Println("s4 and s5 are equals")
	} else {
		fmt.Println("s4 and s5 are not equals")
	}
}

func stringCompareUsingByteSlice() {
	//strings.Compare(s1,s2) is not available only bytes have Compare function
	/*s1 := []string{"AB", "BC", "11", "ZA", "10"}
	s2 := []string{"AB", "BC", "11", "ZA", "00"}
	if strings.Compare(s1, s2) == 0 {
		fmt.Println("s1 and s2 are equals")
	} else {
		fmt.Println("s1 and s2 are not equals")
	}*/

	//using string you can create slice of byte. Then can use Compare(). Since Compare is not in strings
	s4 := []byte("JAVA") //Use () NOT {}
	s5 := []byte("JAVA") //Use () NOT {}
	if bytes.Compare(s4, s5) == 0 {
		fmt.Println("s4 and s5 are equals")
	} else {
		fmt.Println("s4 and s5 are not equals")
	}
}

func byteSclicesEqualityCheck() {
	s1 := []byte{'A', 'B', '1', 'Z', '0'}
	s2 := []byte{'A', 'b', '1', 'Z', '0'}
	s3 := []byte{'A', 'B', '1', 'Z', '0'}

	var result1 bool = bytes.Equal(s1, s2)
	if result1 {
		fmt.Println("s1 and s2 are equals")
	} else {
		fmt.Println("s1 and s2 are not equals")
	}
	if bytes.Equal(s1, s3) {
		fmt.Println("s1 and s3 are equals")
	} else {
		fmt.Println("s1 and s3 are not equals")
	}

	s4 := []byte("Java") //Use () NOT {}
	s5 := []byte("JAVA") //Use () NOT {}
	if bytes.Equal(s4, s5) {
		fmt.Println("s4 and s5 are equals")
	} else {
		fmt.Println("s4 and s5 are not equals")
	}
}

func intSortOperation() {
	slc1 := []int{99, 65, 15, 426, 325, -52}
	slc2 := []int{15, 65, 99, 482}
	fmt.Println("Before: slc1: ", slc1)
	fmt.Println("Before: slc2: ", slc2)

	fmt.Println("sort.IntsAreSorted(slc1): ", sort.IntsAreSorted(slc1))
	fmt.Println("sort.IntsAreSorted(slc2): ", sort.IntsAreSorted(slc2))

	if !sort.IntsAreSorted(slc1) {
		sort.Ints(slc1)
		fmt.Println("After: slc1: ", slc1)
	}
	if !sort.IntsAreSorted(slc2) {
		sort.Ints(slc2)
		fmt.Println("After: slc2: ", slc2)
	}
}

func stringSortOperation() {
	slc1 := []string{"AB", "ACB", "ZA", "ab", "LM", "C"}
	slc2 := []string{"AB", "CD", "EF", "GH"}

	fmt.Println("Before: slc1: ", slc1)
	fmt.Println("Before: slc2: ", slc2)

	fmt.Println("sort.StringsAreSorted(slc1): ", sort.StringsAreSorted(slc1))
	fmt.Println("sort.StringsAreSorted(slc2): ", sort.StringsAreSorted(slc2))

	if !sort.StringsAreSorted(slc1) {
		sort.Strings(slc1)
		fmt.Println("After: slc1: ", slc1)
	}
	if !sort.StringsAreSorted(slc2) {
		sort.Strings(slc2)
		fmt.Println("After: slc2: ", slc2)
	}
}
