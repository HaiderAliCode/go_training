package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//task3
//elegement approach to create server
func createServer(name string, port int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello"+name)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

func main() {
	//to run all this on one core
	// runtime.GOMAXPROCS(1)

	//multiple http server

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "hello"+r.Host)
	// })

	// go func() {
	// 	http.ListenAndServe(":9000", nil)
	// }()

	// http.ListenAndServe(":9001", nil)

	//task 2
	// //syncgroup to spawn multiple servers
	// wg := new(sync.WaitGroup)

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "hello"+r.Host)
	// })

	// go func() {
	// 	http.ListenAndServe(":9000", nil)
	// 	wg.Done()
	// }()
	// wg.Add(1)

	// go func() {
	// 	http.ListenAndServe(":9001", nil)
	// 	wg.Done()
	// }()
	// wg.Add(1)

	// wg.Wait()

	//task3
	// wg := new(sync.WaitGroup)
	// wg.Add(2)

	// go func() {
	// 	server := createServer("server1", 9000)
	// 	fmt.Println(server.ListenAndServe())
	// 	wg.Done()
	// }()

	// go func() {
	// 	server := createServer("server2", 9001)
	// 	fmt.Println(server.ListenAndServe())
	// 	wg.Done()
	// }()

	// wg.Wait()

	/*The Server.SetKeepAliveEnabled method enables the HTTP persistent connection. By default, this is enabled for better performance.
	server.SetKeepAlivesEnabled(false)*/

	//can close when needed
	// error := server.Close()

	// //task 4
	// //closing server wisely
	// exitSignal := make(chan interface{})

	// server := &http.Server{
	// 	Addr:    ":9090",
	// 	Handler: nil,
	// }

	// //will close server after 3 seconds
	// time.AfterFunc(3*time.Second, func() {
	// 	fmt.Println("server closed", server.Close())
	// 	close(exitSignal)
	// })

	// err := server.ListenAndServe()
	// fmt.Println("ListenAndServeError:", err)

	// //blocking waiting for server to close
	// <-exitSignal

	// fmt.Println("Main() exit complete")

	/*he server.Shutdown method behaves exactly like server.Close method but with some added safety. If you want to gracefully close active TCP connections, then Shutdown() is better compared to Close method.*/

	//task 5
	//shutting down server gracefully and cleaning up the resources

	wg := new(sync.WaitGroup)
	wg.Add(2)

	server := &http.Server{
		Addr:    ":9000",
		Handler: nil,
	}

	server.RegisterOnShutdown(func() {
		fmt.Println("closing server running cleanup")
		wg.Done()
	})

	time.AfterFunc(3*time.Second, func() {
		err := server.Shutdown(context.Background())
		fmt.Println("execution of server closing in progress err:", err)
		wg.Done()
	})

	fmt.Println("server listening in progress", server.ListenAndServe())
	wg.Wait()
	fmt.Println("main() exited")

	//can check and handle http.ListenAndServe() call is equal to http.ErrServerClosed to handle error on closing
}
