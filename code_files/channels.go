package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func squareval(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	//always close channel from near its sender dont do it near receiver
	close(c)
}

func printchannelvalues(c chan string) {
	for val := range c {
		fmt.Println(val)
	}
}

func funcforchan1(c chan string) {
	time.Sleep(time.Millisecond * 300)
	c <- "hello"
}

func funcforchan2(c chan string) {
	time.Sleep(time.Millisecond * 200)
	c <- "world"
}

func main() {
	// nil channel
	// var c chan int
	// fmt.Println(c)

	//empty channel which contains a pointer to an memory
	// c1 := make(chan int)
	// fmt.Println(c1)

	//reading from channel
	//this will call a deadlock since there is no other sender to chanel
	// data := <-c1
	// fmt.Println(data)
	// c1 <- 1

	//sending data to channel
	// c2 := make(chan string)
	// go greet(c2)
	// c2 <- "Haider"

	//closing a channel after using
	// close(c2)

	//error of channel being closed
	// c2 <- "test2"

	//if we are trying to read data from a channel and there is no data in channel then it may block it
	//example only main goroutine doing some channel stuff //deadlock error
	// c1 := make(chan string)
	// c1 <- "test"

	// infinite for loop to get data from a channel
	// c2 := make(chan int)
	// go squareval(c2)

	// for {
	// 	val, ok := <-c2
	// 	if ok == false {
	// 		fmt.Println("error occured at val", val)
	// 		break
	// 	} else {
	// 		fmt.Println("got value", val)
	// 	}
	// }

	//easy for loop syntax
	// for val := range c2 {
	// 	fmt.Println(val)
	// }

	//channel with buffer
	// c2 := make(chan string, 3)
	// //channel length and capacity
	// fmt.Println("channel", c2)
	// //length is containing elements in buffer
	// fmt.Println("length", len(c2))
	// //capacity is total number of values a channel can have
	// fmt.Println("capacity", cap(c2))

	// go printchannelvalues(c2)

	// c2 <- "hello"
	// c2 <- "123"
	// c2 <- "456"             //until here goroutine was not blocked because buffer is still in capacity
	// c2 <- "buffer overflow" // here main will block because buffer overflow happened and other goroutine will run

	//unidirectional channels
	//receive only channel
	// roc := make(<-chan string)
	//send only channel
	// soc := make(chan<- string)

	//can convert bidirectional to unidirectional
	//this function's channel will now be readonly
	// func(roc <-chan int){}

	//channels with anonymous goroutine
	// go func(c chan string){
	// 	fmt.Println(<-c)
	// }

	//select statement in golang
	chan1 := make(chan string)
	chan2 := make(chan string)

	go funcforchan1(chan1)
	go funcforchan1(chan2)

	// both statements are blocking channels but when one will hit than it will unblock others
	select {
	case res := <-chan1:
		fmt.Println("chan1 hit", res)
	case res := <-chan2:
		fmt.Println("chan2 hit", res)
	}

	fmt.Println("main() stoppped")
}
