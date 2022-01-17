package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	LastName, Firstname string
}

//a interface's method
func (e Employee) fullname() string {
	return e.Firstname + " " + e.LastName
}

////multiple methods of same name
type emp1 struct {
	Name string
}

type cmp1 struct {
	Name string
}

func (e emp1) Salary() int {
	return 1
}

func (c cmp1) Salary() int {
	return 1000
}

//changing value of an object by reference
func (e *emp1) changename(newname string) {
	(*e).Name = newname
	//or
	e.Name = newname
}

//changing nested stuct of a struct object
type company1 struct {
	Name string
}

type employee3 struct {
	company company1
}

func (e *employee3) changecompany(newcompany string) {
	e.company.Name = newcompany
}

//adding function to existing types
type NewString string

func (s NewString) toupper() string {
	str := string(s)
	return strings.ToUpper(str)
}

func main() {
	emp13 := Employee{"Haider", "Ali"}
	fmt.Println(emp13.fullname())

	emp := emp1{"Haider"}
	cmp := emp1{"Invozone"}

	fmt.Println(emp)
	fmt.Println(cmp)
	emp.changename("Haider Ali")
	fmt.Println(emp)

	emp10 := employee3{
		company: company1{"Programmers Force"},
	}
	fmt.Println(emp10)

	emp10.changecompany("Invozone")
	fmt.Println(emp10)

	newstr := NewString("hello world")
	fmt.Println(newstr.toupper())
}
