package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type MyError struct{}

func (myError *MyError) Error() string {
	return "Something went wrong"
}

func Divide(a int, b int) (int, error) {
	if b == 0 || a == 0 {
		return 0, errors.New("cannot divide value by 0")
	} else {
		return a / b, nil
	}
}

type HttpError struct {
	status int
	method string
}

func (httperr *HttpError) Error() string {
	return fmt.Sprintf("something went wrong with request: ", httperr.method, httperr.status)
}

func GetUserEmail(userId int) (string, error) {
	return "", &HttpError{403, "GET"}
}

func main() {
	//this is how we can create a simple error message but...
	MyErr := &MyError{}
	fmt.Println(MyErr)

	//go provides another interfac3
	err := errors.New("something went wrong")
	fmt.Println(err)

	//or we can also use fmt errorf
	err2 := fmt.Errorf("something went wrong")
	fmt.Println(err2)

	if result, err := Divide(1, 0); err != nil {
		fmt.Println("error occured", err)
	} else {
		fmt.Println(result)
	}

	if email, err := GetUserEmail(1); err != nil {
		// fmt.Printf("%T", err)

		if errVal := err.(*HttpError); errVal.status != 200 {
			fmt.Println("something went wrong with request")
		}
	} else {
		fmt.Println("user email is", email)
	}

	//to get a stacktrace from an error we need to import github util
	//github.com/pkg/error/ -> go mod tidy to install additional depedencies
	originalerr := errors.New("original error")
	newerr := errors.Wrap(originalerr, "additional comment")
	fmt.Println("modified error", newerr)
	fmt.Println("original error", errors.Cause(newerr))
	fmt.Printf("stacked trace error %+v", newerr)

	//type assertion of errors
	// switch err := somefunction(); err.(type) {
	// 	case nil:
	// 		fmt.Println("File successfully saved.")
	// 	case *typeone:
	// 		fmt.Println("Network Error:", err)
	// 	case *typetwo:
	// 		fmt.Println("File save Error:", err)
	// }

	//we can ignore error using a black identifier
	// abc, _ = funcThrowingAnError()

	//tips
	//dont log multiple errors

}
