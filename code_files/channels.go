package main

import (
	"fmt"
	"sync"
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
	// time.Sleep(time.Millisecond * 300)
	c <- "hello"
}

func funcforchan2(c chan string) {
	// time.Sleep(time.Millisecond * 200)
	c <- "world"
}

func funcforwait(c chan string) {
	time.Sleep(time.Second * 3)
	c <- ""
}

func getroutine(wgroup *sync.WaitGroup, elem int) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("service called on instance", elem)
	wgroup.Done()
}

func handleworkerpool(sender <-chan int, receiver chan<- int, instance int) {
	for num := range sender {
		time.Sleep(time.Millisecond)
		fmt.Println("function handler", instance)
		receiver <- num
	}
}

var i int

//implmenting mutex here
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i = i + 1
	m.Unlock()
	wg.Done()
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

	// select statement in golang
	// chan1 := make(chan string)
	// chan2 := make(chan string)

	// go funcforchan1(chan1)
	// go funcforchan1(chan2)

	// // both statements are blocking channels but when one will hit than it will unblock others
	// select {
	// case res := <-chan1:
	// 	fmt.Println("chan1 hit", res)
	// case res := <-chan2:
	// 	fmt.Println("chan2 hit", res)
	// }

	//if we remove sleep from goroutines than you may get result from either of goroutine and this approach can be used in load balancing

	//blocking statement to transfer control from main goroutine to other
	// time.Sleep(time.Millisecond * 10)

	// select {
	// case res := <-chan1:
	// 	fmt.Println("chan1 hit", res)
	// case res := <-chan2:
	// 	fmt.Println("chan2 hit", res)

	//this is a non blocking statement so main goroutine will never block and exist
	//but if we put sleep above select than we might get some data

	// default:
	// 	fmt.Println("nothing received")
	// }

	// //if we use nil channel we might het trouble in select statement
	// var c1 chan string
	// select {
	// case res := <-c1:
	// 	fmt.Println(res)
	// //unless we have default
	// default:
	// 	fmt.Println("no response")
	// }

	//timeout after some time to run default case
	// chan2 := make(chan string)

	// go funcforwait(chan2)

	// select {
	// case res := <-chan2:
	// 	fmt.Println("got result from chan2", res)
	// case <-time.After(time.Second * 2):
	// 	fmt.Println("timeout")
	// }

	//empty select
	//like for{} select statment also runs forever and its a blocking statement unless one of the cases occured otherwise its a deadlock
	// select {}

	//waitgroup
	//when you need to check if all goroutines finished their job
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go getroutine(&wg, i)
	// }
	// wg.Wait()

	//worker pool
	// sender := make(chan int, 10)
	// receiver := make(chan int, 10)

	// for i := 0; i < 3; i++ {
	// 	go handleworkerpool(sender, receiver, i)
	// }

	// for i := 0; i < 5; i++ {
	// 	sender <- i * 2
	// }

	// fmt.Println("sent 5 items in sender")

	// close(sender)

	// for i := 0; i < 5; i++ {
	// 	results := <-receiver
	// 	fmt.Println(results)
	// }

	// /////mutex
	// var wg sync.WaitGroup
	// var m sync.Mutex

	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go worker(&wg, &m)
	// }

	// wg.Wait()
	// // race condition is when multiple goroutines tries to access open variable and we see unsual results
	// fmt.Println("value of i is", i)

	fmt.Println("main() stoppped")
}
