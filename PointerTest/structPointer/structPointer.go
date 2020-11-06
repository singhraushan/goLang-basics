package main

import "fmt"

func main() {
	// you can decllare struct inside fun also
	type employee struct {
		name string
		id   int
	}
	emp := employee{"Raushan", 1}
	fmt.Println("emp.name: ", emp.name) //Raushan
	var pt *employee = &emp
	fmt.Println("(*pt).name: ", (*pt).name) //Raushan
	fmt.Println("pt.name: ", pt.name)       //Raushan
	fmt.Println("after pt: ", *pt)
	pt.name = "Amit"
	fmt.Println("after pt: ", *pt)
	fmt.Println("after emp: ", emp)
}
