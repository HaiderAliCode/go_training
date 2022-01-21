package main

import "fmt"

func accessElement(a []int, index int) int {
	if len(a) > index {
		return a[index]
	} else {
		panic("out of bound of array")
	}
}

func defMain() {
	fmt.Println("defMain hit")
}

func defFoo() {
	//recovering using recover function
	if r := recover(); r != nil {
		fmt.Println("psychiatrist is here for your assistance!")
	}
	fmt.Println("defFoo hit")
}

func endFoo() {
	fmt.Println("endFoo start")
}

func foo() {
	fmt.Println("got in foo function")

	defer defFoo()

	panic("i am panicing, help me")

	defer endFoo()

	fmt.Println("exiting foo function")
}

func recovermagic(a []int, index int) (v int) {
	fmt.Println("hello world")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("cant do anything here")
			//return last element if there is an error
			v = a[len(a)-1]
		}
	}()

	v = a[index]
	return
}
func main() {

	// slice := []int{1, 2, 3, 4}
	// fmt.Println("access second element", accessElement(slice, 3))
	// fmt.Println("access second element", accessElement(slice, 4))

	//when a panic occurs go doesnt terminate program immediately it do all the unstacking to reach main function
	//and then it terminates though no new statement runs its just all going back.

	//all those functions which are in defer stack runs in lifo order
	// defer defMain()

	// foo()

	// fmt.Println("Main ended")

	//recovering from a panic
	//we can recover a program from exiting using recover function
	// defer defMain()

	// foo()

	// //see our program ran succesfully
	// fmt.Println("Main ended")

	//we can do many things with recover function
	temp_slice := []int{1, 2, 3}
	fmt.Println(recovermagic(temp_slice, 10))

}
