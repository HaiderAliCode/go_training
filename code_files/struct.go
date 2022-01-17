package main

import "fmt"

type Company struct {
	name string
}

//initializing struct
type Employee2 struct {
	name     string
	salary   int
	fulltime bool
	company  Company //nested struct
}

//Anoymous Company field
type Employee3 struct {
	name     string
	salary   int
	fulltime bool
	Company  //nested struct
}

//initializing struct
type Employee struct {
	name     string
	salary   int
	fulltime bool
}

//anonymous fields struct
type anonstruct struct {
	string
	int
}

//json with metadata
type Employee6 struct {
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	salary    int    `json: "salary"`
	fullTime  int    `json: "fullTime"`
}

func main() {
	var emp1 Employee
	//printing empty object
	fmt.Println(emp1)
	//setting struct value
	emp1.name = "Haider"
	//getting struct value
	fmt.Println(emp1.name)

	//creating struct using shorthanded notation
	emp2 := Employee{
		name:     "Haider Ali",
		salary:   190000,
		fulltime: true,
	}
	fmt.Println(emp2)

	//third way of initializing a struct
	emp3 := Employee{"Haider", 20000, true}
	fmt.Println(emp3)

	//anonymous struct
	emp4 := struct {
		name string
	}{
		name: "ali",
	}
	fmt.Println(emp4)

	//pointer to a struct
	emp5 := &Employee{
		name: "haider",
	}
	//with reference
	fmt.Println(emp5)
	//without reference
	fmt.Println(*emp5)
	fmt.Println((*emp5).name)

	//anonymous struct
	anonstruct1 := anonstruct{"haider", 1}
	fmt.Println(anonstruct1)

	//setting value of nested struct
	emp6 := Employee2{
		company: Company{"InvoZone"},
	}
	fmt.Println(emp6)

	//field promotion with anonymous struct
	emp7 := Employee3{
		Company: Company{"Invozone"},
	}
	fmt.Println(emp7.Company)

}
