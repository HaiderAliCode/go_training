package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var signal = make(chan bool)

var c = make(chan int)

func square(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return //kill goroutine
		case c <- i * i:
			i++
		}
	}
}

func printNumbers() {

	counter := 1
	for {
		select {
		//stopper with public variable
		case <-signal:
			return
		default:
			time.Sleep(100 * time.Millisecond)
			counter++
			fmt.Println(counter)
		}
	}
}

func main() {

	// go printNumbers()
	// fmt.Println("main() started", runtime.NumGoroutine())
	// time.Sleep(time.Second)
	// signal <- true

	// fmt.Println("after sending stop signal", runtime.NumGoroutine())
	// fmt.Println("", runtime.NumGoroutine())

	//normal context
	ctx, cancel := context.WithCancel(context.Background())

	// //timeline context
	// deadline := time.Now().Add(3 * time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), deadline)

	go square(ctx)
	//getting 5 squares
	for i := 0; i < 5; i++ {
		fmt.Println("value of square is", <-c)
	}

	cancel()
	//defer cancel if main exists before deadine
	// defer cancel()

	time.Sleep(time.Second * 10)
	//we get 1 couroutine before ending main
	//we can see that above goroutine was still running even though we dont need that anymore
	fmt.Println("number of goroutines at the end of main()", runtime.NumGoroutine())
}
