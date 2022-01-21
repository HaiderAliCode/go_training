package unittesting

import "fmt"

// simple syntax of testing
// import "testing"
// func TestAbc(t *testing.T) {
//     t.Error() // to indicate test failed
// }

//can create a unittesting cache with command
// go test ./unittesting(folder name) -c
//and run the test with
// ./unittesting.test -test.v
//it can save you time
//dont overwrite cache i there is no change in your file e.g. variable name change not gonna change the output or flow

//you can create a tessdata directory inside your package and put your file in there it wont be used during build
//and you can also use the test data from it

func add(hello string) string {
	if len(hello) == 0 {
		return "Hello Dude!"
	} else {
		return fmt.Sprintf("Hello %v!", hello)
	}
}

func squareSlice(s []int) []int {
	squareS := make([]int, len(s))

	for index, val := range s {
		squareS[index] = val * val
	}

	return squareS
}
