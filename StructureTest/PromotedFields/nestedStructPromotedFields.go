package main

import "fmt"

func main() {
	uniqueFieldName()
	fmt.Println()
	duplicateFieldName()
}

func duplicateFieldName() {
	type address struct {
		name, city, state string
		pincode           int
	}
	type student struct {
		name    string
		age     int
		address // anonymous field
	}

	s := student{
		"Raushan",
		30,
		address{
			"duplicateFieldName", "patna", "bihar", 848204,
		},
	}
	//all are promoted fields. so no need to use like s.address.<filedName>
	fmt.Println(" duplicate field name s.name: ", s.name) // highest priority to parent struct field name
	fmt.Println(" duplicate field name s.address.name: ", s.address.name)
	fmt.Println("s.city: ", s.city)
	fmt.Println("s.state: ", s.state)
	fmt.Println("s.pincode: ", s.pincode)
	fmt.Println("s.address.pincode: ", s.address.pincode) // same like s.pincode

	fmt.Println("s.name: ", s.name)
	fmt.Println("s.age: ", s.age)
}

func uniqueFieldName() {
	type address struct {
		city, state string
		pincode     int
	}
	type student struct {
		name    string
		age     int
		address // anonymous field
	}

	s := student{
		"Raushan",
		30,
		address{
			"patna", "bihar", 848204,
		},
	}
	//all are promoted fields. so no need to use like s.address.<filedName>
	fmt.Println("s.city: ", s.city)
	fmt.Println("s.state: ", s.state)
	fmt.Println("s.pincode: ", s.pincode)
	fmt.Println("s.address.pincode: ", s.address.pincode) // same like s.pincode

	fmt.Println("s.name: ", s.name)
	fmt.Println("s.age: ", s.age)
}
