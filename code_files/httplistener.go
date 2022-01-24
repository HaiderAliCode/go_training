package main

import (
	"fmt"
	"io"
	"net/http"
)

type httpHandler struct{}

func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello ")
	fmt.Fprint(res, "world! ")
	res.Write([]byte("❤️"))
}

// type homeHandler struct{}

// func (h homeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "You are in home handler")
// }

// type golangHandler struct{}

// func (h golangHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "you are in golang page")
// }

func main() {
	// handler := httpHandler{}
	// http.ListenAndServe(":8080", handler)

	//working with mutex
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { io.WriteString(rw, "you are in home page") })
	// mux.HandleFunc("/home/golang", func(rw http.ResponseWriter, r *http.Request) { io.WriteString(rw, "you are in golang page") })
	// http.ListenAndServe(":8080", mux)

	//working without mutex
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { io.WriteString(rw, "you are in home page") })
	// http.HandleFunc("/home/lang", func(rw http.ResponseWriter, r *http.Request) { io.WriteString(rw, "you are in golang page") })
	// http.ListenAndServe(":8080", nil)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		header := rw.Header()

		header.Set("Content-Type", "application/json")
		header.Set("Date", "01/01/2020")
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, `{"status":"failue"}`)

	})
	http.ListenAndServe(":8080", nil)

}
