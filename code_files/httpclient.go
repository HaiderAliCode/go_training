package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	//default http client
	// var defaultClient = &http.Client{}

	// res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// fmt.Printf("%s\n", data)

	//getting a file from internet
	// path := "https://bit.ly/2IRnmVm"

	// res, err := http.Get(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()

	// //can get header from image and set it back etc
	// if res.Header.Get("Content-Type") == "image/jpeg" {
	// 	ioutil.WriteFile("httpimage.jpeg", data, 0777)
	// 	fmt.Println("image downloaded successfully")
	// } else {
	// 	fmt.Println("response is not an image")
	// }

	//http post request
	// requestBody := strings.NewReader(`
	// {
	// 	"name":"haider",
	// 	"age" : "25"
	// }
	// `)

	// res, err := http.Post("http://dummy.restapiexample.com/api/v1/create", "application/json", requestBody)
	// if err != nil {
	// 	log.Fatal("error occured while processing post request")
	// }

	// data, _ := ioutil.ReadAll(res.Body)

	// res.Body.Close()
	// requestContentType := res.Request.Header.Get("Content-Type")
	// fmt.Println("received content type is", requestContentType)

	// fmt.Printf("%s\n", data)

	//custom http request
	// reqUrl, _ := url.Parse("http://dummy.restapiexample.com/api/v1/create")
	// reqBody := ioutil.NopCloser(strings.NewReader(`
	// {
	// 	"name":"haider",
	// 	"age" : "25"
	// }
	// `))
	// req := &http.Request{
	// 	Method: "POST",
	// 	URL:    reqUrl,
	// 	Header: map[string][]string{
	// 		"Content-Type": {"application/json; charset=UTF-8"},
	// 	},
	// 	Body: reqBody,
	// }
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Fatal("error occured while processing request")
	// }
	// data, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// fmt.Printf("%s\n", data)

	//or we can simplify things using http.newrequest
	// reqBody := strings.NewReader(`{"name": ""haider}`)
	// req, err := http.NewRequest("POST", "http://dummy.restapiexample.com/api/v1/create", reqBody)

	// if err != nil {
	// 	log.Fatal("error occured while processing request")
	// }

	// req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// res, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// data, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// fmt.Printf("%s\n", data)

	//http with timeout
	client := &http.Client{
		Timeout: 1000 * time.Millisecond,
	}

	res, err := client.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		urlErr := err.(*url.Error)
		if urlErr.Timeout() {
			fmt.Println("timeout error occured while processing request")
		}

		log.Fatal("Error: ", err)

	} else {
		fmt.Println("Success: status code:", res.StatusCode)
	}

}
