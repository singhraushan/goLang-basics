package main

import "fmt"

type Address struct {
	name, street, city, state string
	pincode                   int
}

type Address2 struct {
	name    string
	street  string
	city    string
	state   string
	pincode int
}

var a Address
var a2 Address2

func main() {
	fmt.Println("Default value of Address:", a)   // default value would be { ,0}
	fmt.Println("Default value of Address2:", a2) // default value would be { ,0}

	a3 := Address{"Raushan", "korai", "begusrai", "bihar", 848204}
	fmt.Println("Address a3:", a3) //Address a3: {Raushan korai begusrai bihar 848204}

	a4 := Address{name: "Raushan", street: "korai", city: "begusrai", state: "bihar",
		pincode: 848204, //if line changes then comma needed else not required
	}
	fmt.Println("Address a4:", a4) //Address a4: {Raushan korai begusrai bihar 848204}

	//accessing field
	fmt.Println("Address a4.name:", a4.name)
	fmt.Println("Address a4.pincode:", a4.pincode)

	a4.name = "Terence"
	fmt.Println("Address a4.name:", a4.name)

}
