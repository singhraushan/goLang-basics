package main

import "fmt"

type author struct {
	name       string
	department string
	salary     int
}

func (a author) show() {
	fmt.Println("name : ", a.name)
	fmt.Println("department : ", a.department)
	fmt.Println("salary : ", a.salary)
}

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {

	a := author{
		name:       "Raushan",
		department: "CSE",
		salary:     2500,
	}

	a.show()
	(&a).show() // this also work

	value1 := data(10)
	value2 := data(20)

	fmt.Println("value1.multiply", value1.multiply(value2))
	fmt.Println("main end")

}
