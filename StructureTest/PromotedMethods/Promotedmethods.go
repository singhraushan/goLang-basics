package main

import "fmt"

type student struct {
	name    string
	age     int
	address // anonymous field
}
type address struct {
	city, state string
	pincode     int
}

func (a address) promotedMethod(salary int) (x int) {
	return a.pincode * salary
}

func (s student) promotedMethod(salary int) int {
	return s.age * salary
}

func main() {
	promotedMethodTest()
}

func promotedMethodTest() {
	s := student{
		"Raushan",
		30,
		address{
			"patna", "bihar", 848204,
		},
	}
	fmt.Println(s.promotedMethod(5))         //call (a address) promotedMethod() if (a address) promotedMethod() not available
	fmt.Println(s.address.promotedMethod(5)) // always call child method
}
