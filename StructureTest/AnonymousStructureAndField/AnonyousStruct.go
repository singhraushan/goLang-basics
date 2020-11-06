package main

import "fmt"

func main() {

	declareAndDefineAnonymousStruct()
	declareAndDefineAnonymousStructField()
}
func declareAndDefineAnonymousStruct() {
	s := struct {
		name  string
		id    int
		money float64
	}{ //need to provide value to every field
		"amit",
		15,
		158.6}

	fmt.Print("s:", s, "\n")
}

func declareAndDefineAnonymousStructField() {
	type employee struct {
		int
		float64
	}

	type employee2 struct {
		int
		//int  C.E.
	}
	e := employee{
		15,
		20,
	}
	fmt.Println("e:", e)

	type employee3 struct {
		name string
		id   int
		int
		float64
	}
	e2 := employee3{
		"Raushan", 15,
		25,
		20.5,
	}
	fmt.Print("e2:", e2, "\n")
	fmt.Print(e2.int)
	//here you need to pass value to all field without mentioning name
	/*e2 := employee3{
		name: "Raushan",
		id:   15,
		25,
		20.5,
	}*/
}
