package main

import "fmt"

type student struct {
	name   string
	branch string
}

type teacher struct {
	language string
	marks    int
}

func (s student) modify(bname string) {
	s.branch = bname
}

func (t teacher) modify(language string) {
	t.language = language
}

func main() {
	t := teacher{
		marks: 65,
	}

	s := student{
		name: "aarti",
	}
	fmt.Println("teacher language before:", t.language) //deault string value
	t.modify("Hindi")
	fmt.Println("teacher language after:", t.language) // not chaged at calling place becz not pointer

	fmt.Println("student language before:", s.name)
	s.modify("sangeeta")
	fmt.Println("student name after:", s.name) // not chaged at calling place becz not pointer
}
