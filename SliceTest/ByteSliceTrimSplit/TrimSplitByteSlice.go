package main

import (
	"bytes"
	"fmt"
)

func main() {
	trimTest()
	fmt.Println("\n======================")
	splitTest()
}
func trimTest() {
	slice1 := []byte{'A', 'B', '1', 'Z', '0'}
	slice2 := []byte("****Welcome, to, GeeksforGeeks****")
	slice3 := []byte{'%', 'g', '%', 'e', '%', 'e', '%', 'k', '%', 's', '%'}
	slice4 := []byte("%%Java***")

	fmt.Println("Original Slice:")
	fmt.Printf("Slice 1: %s", slice1) //Printf need format type(like %s, %d..) if more than one value passing
	fmt.Printf("\nSlice 2: %s", slice2)
	fmt.Printf("\nSlice 3: %s", slice3)
	fmt.Printf("\nSlice 4 %s: ", slice4) // youu need to give %s else print asci value

	res01 := bytes.Trim(slice1, "01") // not remove 1. becuse 1 is not in leading or/and trailing
	// Display the results
	fmt.Printf("\nNew Slice:\n")
	fmt.Printf("\nSlice 1: %s", res01)
	fmt.Printf("\nSlice 4: %s", bytes.Trim(slice4, "%*"))

	// Creating and trimming
	// the slice of bytes
	// Using Trim function
	res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	res3 := bytes.Trim([]byte("^^Geek&&"), "$")
	str3 := string(res3) // covert byte to string
	// Display the results
	fmt.Printf("Final Slice:\n")
	fmt.Printf("\nSlice 1: %s", res1) // to print byte value you need to give %s else print asci for each character
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)
	fmt.Println("\nSlice str 3:", str3)
}

func splitTest() {
	sl1 := []byte{'A', 'p', 'p', 'l', 'e'}
	sl2 := []byte("Java")
	sl3 := []byte{}
	//syntax: bytes.Split(inputbyte []byte seperation []byte) [][]byte
	fmt.Printf("slice 1: %s\n", bytes.Split(sl1, []byte("")))   //[A p p l e]
	fmt.Printf("slice 2: %s\n", bytes.Split(sl2, []byte{'a'}))  //[J v ]
	fmt.Printf("slice 2: %s\n", bytes.Split(sl2, []byte("va"))) //[Ja ]
	fmt.Printf("slice 3: %s\n", bytes.Split(sl3, []byte(",")))  //[]

}
