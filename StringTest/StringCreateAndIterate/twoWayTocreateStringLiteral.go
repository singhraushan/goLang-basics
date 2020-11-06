package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("=======createTwoWayString======")
	createTwoWayString()
	fmt.Println("=======tryToModifyString======")
	tryToModifyString()
	fmt.Println("=======iterateOverString======")
	iterateOverString()
	fmt.Println("=======stringFromSlice======")
	stringFromSlice()
	fmt.Println("=======stringLenCal======")
	stringLenCal()
	fmt.Println("=======twoWaysToCompareStrings=====")
	twoWaysToCompareStrings()
}
func createTwoWayString() {
	//using double quotes("")
	str1 := "Hello,\nI'm Raushan" //this suport multi line
	//using backticks(``)
	str2 := `Hello,\nI'm Raushan` //this not support multi line
	fmt.Println("str1:", str1)    //o/p:- two line only
	fmt.Println("str2:", str2)    //o/p:- one line only
}

// string are immutable like java
func tryToModifyString() {
	str1 := "YOu can't modify string"
	fmt.Println("str1:", str1)

	//C.E. can't assign to string
	//str1[0] = 'J'
	//str1[0] = "J"
}

func iterateOverString() {
	for index, char := range "Hello" {
		fmt.Printf("Index of %c is %d\n", char, index)
	}

	//accssing individual bytes of string
	fmt.Println("---------------------")
	var str1 string = "Amit"
	for c := 0; c < len(str1); c++ {
		fmt.Printf("\nCharacter = %c Bytes = %v", str1[c], str1[c])
	}
}

func stringFromSlice() {
	slice1 := []byte{0x47, 0x65, 'e', 'k', 's'}
	slice2 := []byte{0x0047, 0x0065, 'e', 'k', 's'}
	slice3 := []byte("Geeks")

	str1 := string(slice1)
	str2 := string(slice2)
	str3 := string(slice3)

	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
	fmt.Println("str3: ", str3)
}

func stringLenCal() {
	str1 := "Hello,\nI'm Raushan"
	str2 := `Hello,\nI'm Raushan`
	str3 := "Java"

	fmt.Println("len of str1: ", len(str1)) //18
	fmt.Println("len of str2: ", len(str2)) //19
	fmt.Println("len of str3: ", len(str3)) // 4

	fmt.Println("RuneCountInString(i.e. len) of str1: ", utf8.RuneCountInString(str1)) //18
	fmt.Println("RuneCountInString(i.e. len) of str2: ", utf8.RuneCountInString(str2)) //19
	fmt.Println("RuneCountInString(i.e. len) of str3: ", utf8.RuneCountInString(str3)) //4
}

func twoWaysToCompareStrings() {
	str1 := "Hello"
	str2 := `Hello` //uing backticks
	str3 := "Hi"

	fmt.Println("str1 && str2 are equals :", str1 == str2)
	fmt.Println("str1 && str3 are equals :", str1 == str3)
	fmt.Println("str2 && str3 are equals :", str2 == str3)

	fmt.Println("str1 > str2 :", str1 > str2)
	fmt.Println("str1 > str3  :", str1 > str3)
	fmt.Println("str2 > str3 :", str2 > str3)
	fmt.Println("str1 < str2 :", str1 < str2)
	fmt.Println("str1 < str3  :", str1 < str3)
	fmt.Println("str2 < str3 :", str2 < str3)

	if strings.Compare(str1, str2) == 0 {
		fmt.Println("str1 && str2 are equal.")
	} else {
		fmt.Println("str1 && str2 are NOT equal.")
	}

	if strings.Compare(str1, str3) == 0 {
		fmt.Println("str1 && str3 are equal.")
	} else {
		fmt.Println("str1 && str3 are NOT equal.")
	}

	if strings.Compare(str2, str3) == 0 {
		fmt.Println("str2 && str3 are equal.")
	} else {
		fmt.Println("str2 && str3 are NOT equal.")
	}

}
