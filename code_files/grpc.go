package main

import (
	"fmt"
	"go_code/common"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	//server implementation
	mit := common.NewCollege()

	rpc.Register(mit)
	rpc.HandleHTTP()
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "saying hello from grpc")
	})

	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
	}()

	time.Sleep(time.Second * 1)

	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000")
	var john common.Student = common.Student{ID: 1}

	if err := client.Call("College.Get", 1, &john); err != nil {
		fmt.Println("Error:1 College.Get()", err)
	} else {
		fmt.Printf("data found ", john.FullName())
	}

	if err := client.Call("College.Add", common.Student{
		ID:        1,
		FirstName: "Haider",
		LastName:  "Ali",
	}, &john); err != nil {
		fmt.Println("Error 2: College.Add()", err)
	} else {
		fmt.Println("successfully created\n", john.FullName())
	}

}
