package main

func changepointervalue(p *int) {
	*p = 1
}

func compositepointerhandler(p *[3]int) {
	////can use both ways to access data of composite type model
	p[1] = 9
	(*p)[2] = 8
}

func main() {
	////using ampersand sign to access memory address of a variable
	// var a string = "hello"
	// fmt.Println(a)
	// fmt.Println(&a)

	////we can store a hexadecimal number in a variable but the variable wont know what to do with that address looking number
	////thus proper way to define a variable is var variable_name *Type
	////example
	// a := 1
	// var ab *int
	// ab = &a
	// fmt.Println(ab)

	////or shorthand way of creating is
	// b := 1
	// ba := &b
	// fmt.Println(ba)
	////we can use * with variable to get value of the referencing variable
	// fmt.Println(*ba)
	////changing value using pointer
	// *ba = 2
	// fmt.Println(b)

	////NOTE: The difference between a variable and pointer is that a variable stores the value at a memory address and pointer points to a memory address.

	////passing pointer as an argument
	// a := 2
	// fmt.Println(a)
	// changepointervalue(&a)
	// fmt.Println(a)

	////pointer of composite type
	// a := [3]int{1, 2, 3}
	// fmt.Println(a)
	// compositepointerhandler(&a)
	// fmt.Println(a)

}
