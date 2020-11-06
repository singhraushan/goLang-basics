package main

import "fmt"

type teacher struct {
	name    string
	subject string
	exp     int
	details student
}

type student struct {
	name   string
	branch string
}

var t = teacher{
	name:    "Ramakant",
	subject: "Math",
	exp:     12,
	details: student{
		name:   "Raushan",
		branch: "CSE",
	},
}

func main() {
	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", t.name)
	fmt.Println("Subject: ", t.subject)
	fmt.Println("Experience: ", t.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", t.details.name)
	fmt.Println("Student's branch name: ", t.details.branch)
	//fmt.Println("Student's branch name: ", t.branch) // C.E. if try like promoted even it's part of named struct
}
