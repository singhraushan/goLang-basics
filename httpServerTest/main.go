package main

import (
	"fmt"
	"log"
"net/http"
"encoding/json"
)

type Employee struct{
	 Id int `json:"Id"`
	 Name string `json:"Name"`
	 Salary float64 `json:"Salary"`
}

func main(){
	handleRequests()
}

func handleRequests(){
	log.Println("log:Server starting on: http://localhost:8080")
	http.HandleFunc("/",homePage)
	http.HandleFunc("/employeeList",employeeList)
	log.Fatal(http.ListenAndServe(":8080",nil))	
}

func homePage(w http.ResponseWriter, r *http.Request){
	   fmt.Fprintf(w,"This is Home Page.") 
	   fmt.Println("home page hitted.")
}

func employeeList(w http.ResponseWriter, r *http.Request){
	employees := []Employee{
		Employee{1,"Raushan",70253.78}, Employee{2,"Amit",180253.78},
	}
	json.NewEncoder(w).Encode(employees)
	fmt.Println("employeeList page hitted.")
}