package main

import (
	"fmt"
)

func main() {
	ExpressionSwitch()
	ExpressionSwitch2()
	StringSwitch()
	typeSwitch()
}
func ExpressionSwitch() {

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day") // no need to mention case keyword else C.E.
	}
}

func ExpressionSwitch2() {
	var v int = 2
	switch {
	case v == 1:
		fmt.Println("Hello")
	case v == 2:
		fmt.Println("Bonjour")
	case v == 3:
		fmt.Println("Namstey")
	default:
		fmt.Println("Invalid")
	}
}

// Switch statement without default statement
// Multiple values in case statement
func StringSwitch() {
	var v string = "five"
	switch v {
	case "one":
		fmt.Println("C#")
	case "two":
		fmt.Println("Go")
	case "three", "four", "five":
		fmt.Println("Java")
	}
}

func typeSwitch() {
	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}
}
