package main

import (
	"fmt"
	"time"
)

//go routines are function called with go keyword
func printhello() {
	//if we sleep more than main thread program will just be killed by main thread
	time.Sleep(15 * time.Millisecond)
	fmt.Println("hello world")
}

func main() {
	fmt.Println("first line")

	go printhello()
	// if before line is uncomment then above function will not run
	time.Sleep(10 * time.Millisecond)
	//because main goroutine is still running and when is is paused then other goroutines run
	//other go routines must run within the sleeping time of main thread

	//anonymous goroutines
	go func() {
		fmt.Println("hello anonymous go routine")
	}()

	time.Sleep(6 * time.Millisecond)

	fmt.Println("main ended")
}
