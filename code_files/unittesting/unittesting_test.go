package unittesting

import (
	"reflect"
	"testing"
)

//we can install go get -u github.com/rakyll/gotest
//gotest -v -run TestU //this creates a like function search

// go test -cover shows us coverage of our test run
//go test -coverprofile=cover.txt writes cover profile in cover.txt
//we can use go tool cover -html=cover.txt -o cover.html

func TestUnittest(t *testing.T) {
	myres := add("")

	if myres != "Hello Dude!" {
		t.Error("Failed")
	} else {
		t.Log("Hey bro i am your verbose")
	}
}

func TestSquareSlice(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4}
	resSlice := []int{1, 4, 9, 16}

	result := squareSlice(inputSlice)

	if reflect.DeepEqual(resSlice, result) {
		t.Log("both are equal")
	} else {
		t.Error("SquareSlice, Failed")
	}

}
