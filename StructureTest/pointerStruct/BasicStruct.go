package main

import "fmt"

type Address struct {
	name, street, city, state string
	pincode                   int
}

func main() {
	//(*a3).name && a4.name gives same output

	a3 := &Address{name: "Raushan"}
	fmt.Println("Address (*a3).name:", (*a3).name) //Raushan

	a4 := &Address{name: "Raushan"}
	fmt.Println("Address a4.name:", a4.name) //Raushan

}
