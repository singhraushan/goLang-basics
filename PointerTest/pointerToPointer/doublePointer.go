package main

import "fmt"

func main() {
	var v int = 13
	var pt1 *int = &v
	var pt2 **int = &pt1

	fmt.Println(**pt2)
	**pt2 = 500
	fmt.Println(**pt2) //500
	fmt.Println(*pt1)  //500
	fmt.Println(v)     //500
}
