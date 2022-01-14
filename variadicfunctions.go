package main

import "fmt"

////vardiac are those functions which can take any number of parameters or variable number of parameters

func variadicfunctions(args ...int) {
	fmt.Println(args)
}

////vardiac function multiple parameters, unpacking paramter argument must be provide at the end of parameter list
func variadicfunctions2(param1 string, args ...int) {
	fmt.Println(param1, "===>", args)
}

func main() {
	variadicfunctions(1, 2, 3, 4, 5, 6)

	////throws error if we pass a wrong type
	////variadicfunctions2(1, 2, 3, 4, 5, 6)

	variadicfunctions2("hello", 2, 3, 4, 5, 6)

	////can unpack a slice into it
	s1 := []int{1, 2, 3, 4, 5, 6}
	variadicfunctions2("a slice here", s1...)
}
