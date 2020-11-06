package main

import "fmt"

type author struct {
	name       string
	department string
	salary     int
}

func (a *author) modify(branchname string) {
	(*a).department = branchname
}

func (a *author) show_1(name string) {
	(*a).name = name
}

func (a author) show_2() {
	a.name = "Terence"
	fmt.Println("Author's name(Before) : ", a.name)
}

func main() {

	a := author{
		name:       "amit",
		department: "CSE",
	}
	fmt.Println("department name before:", a.department)
	a.modify("ECE") // calling using value
	fmt.Println("department name after1:", a.department)

	(&a).modify("EE") // calling using address
	fmt.Println("department name after2:", a.department)

	fmt.Println("-------------------------")
	// Calling the show_1 method
	// (pointer method) with value
	a.show_1("Ankit")
	fmt.Println("name Name(After): ", a.name)

	// Calling the show_2 method
	// (value method) with a pointer
	(&a).show_2()
	fmt.Println("Author's name(After): ", a.name)
}
