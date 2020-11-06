package main

import "fmt"

type tank interface {
	Tarea() float64 // no func
	Volume() float64
}

type myValue struct {
	height, radius float64
}

func (m myValue) Tarea() float64 {
	return 2*m.height*m.height + 2*3.14*m.radius
}

func (m myValue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank

	fmt.Println("Default value of interface tank a:", t)    //nil
	fmt.Printf("Default type of interface tank a %T\n:", t) //nil

	t = myValue{10, 14}
	fmt.Println("After initialization: value of interface tank a:", t)    //{10,14}
	fmt.Printf("After initialization: type of interface tank a %T\n:", t) //main.myValue
	fmt.Println("Area of tank:", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())
	fmt.Println("==========================")
	typeAssertions()
	fmt.Println("==========================")
	multiInterfaceImplementedBySingleStruct()
}
func typeAssertions() {
	var v1 interface{} = "Hello"
	fmt.Println("typeAssertions:", v1.(string))

	//fmt.Println("typeAssertions:", v1.(int))// run time error

	// here no error, because found 2c variable and if true then assign lse leave another variable(default value assigned)
	value, isOkay := v1.(int)
	fmt.Println("typeAssertions fordifferent type is :", isOkay, " and value is ", value)

	myFunc("HI")
	myFunc(12)
	myFunc(true)
}

func myFunc(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case float64:
		fmt.Println("Type: float64, Value:", a.(float64))
	case bool:
		fmt.Println("Type: bool, Value:", a.(bool))
	default:
		fmt.Println("Type: string, Value:", a.(string))
	}
}

type authorDetails interface {
	details()
}

type authorArticle interface {
	article()
}
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

func (a author) details() {
	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}
func (a author) article() {
	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

func multiInterfaceImplementedBySingleStruct() {
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the method
	// of the interface 1
	var i1 authorDetails = values
	i1.details()

	// Accessing the method
	// of the interface 2
	var i2 authorArticle = values
	i2.article()
}
