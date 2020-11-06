package main

import "fmt"

type calculateSalary func(pay, marticle int) int
type Author struct {
	name, language string
	pay, marticle  int
	salary         calculateSalary
}

func main() {
	result := Author{
		name:     "Raushan",
		language: "Java",
		pay:      1500,
		marticle: 120,
		salary: func(pay, marticle int) int {
			return pay * marticle // consider local pay & marticle variable value
		},
	}

	// Display values
	fmt.Println("Author's Name: ", result.name)
	fmt.Println("Language: ", result.language)
	fmt.Println("Total number of articles published in May: ", result.marticle)
	fmt.Println("Per article pay: ", result.pay)
	fmt.Println("Total salary: ", result.salary(result.marticle, result.pay))

}
