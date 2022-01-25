package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

"openssl req  -new  -newkey rsa:2048  -nodes  -keyout localhost.key  -out localhost.csr"
"openssl  x509  -req  -days 365  -in localhost.csr  -signkey localhost.key  -out localhost.crt"

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello world")
	// })

	// log.Fatal(http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil)

	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")

	s := &http.Server{
		Addr:    ":9000",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello secure custom world")
	})

	log.Fatal(s.ListenAndServeTLS("", ""))
}
