package main

import (
	"fmt"
	"go_code/models"
)

type fired interface {
	isfired() bool
}

type Company struct {
	Name string
}

//Company struct implementing the interface
func (c Company) isfired() bool {
	return false
}

type Employee struct {
	Name string
	fired
}

type FullNameType func(string, string) string
type Employee2 struct {
	Firstname, Lastname string
	Fullname            FullNameType
}

func main() {

	emp1 := Employee{
		Name:  "Haider Ali",
		fired: Company{"InvoZone"},
	}
	//field promoted because of anonymous interface
	fmt.Println(emp1.isfired())

	//importting go package
	emp3 := models.Employee2{
		Name: "Haider",
	}
	fmt.Println(emp3)

	//function in struct
	emp4 := Employee2{
		Firstname: "Haider",
		Lastname:  "Ali",
		Fullname: func(firstname string, lastname string) string {
			return firstname + " " + lastname
		},
	}
	fmt.Println(emp4)
	fmt.Println(emp4.Fullname(emp4.Firstname, emp4.Lastname))

	//if same struct have two objects than you can compare values with == struct1 == struct2
	//condition is it must have fields which could be comparable

}
