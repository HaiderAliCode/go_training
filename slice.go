package main

////slice is a struct
//// type slice struct {
////     zerothElement *type
////     len int
////     cap int
//// }

func main() {
	////defining empty slice
	// var s []int

	////can create slice slicing through an array
	// fmt.Println("\n!!--array to slice--!!")
	// arr := [...]int{1, 2, 3, 4, 5, 6}
	// arrslice := arr[2:5]
	// fmt.Println("array", arr)
	// fmt.Println("slice", arrslice)

	////length of slice
	// fmt.Println("\n!!--length of slice--!!")
	// cities := []string{
	// 	"Lahore",
	// }
	// fmt.Println("cities", cities)
	// fmt.Println("length of cities is", len(cities))
	// countries := make([]string, 42)
	// fmt.Println("countries", countries)
	// fmt.Println("length of countries is", len(countries))

	////length vs capacity
	// fmt.Println("\n!!--length vs capacity--!!")
	// fmt.Println("array", arr)
	// fmt.Println("slice", arrslice)
	////len of window of slice
	// fmt.Println("length of slice window pointing to", len(arrslice))
	////len of elements in array starting from zeroth element of slice
	// fmt.Println("length of array starting from zeroth element of slice", cap(arrslice))

	////appending slice to slice
	////we did not referenced a slice to an array in this case GO will create an anonymous array
	// s1 := []int{1, 2, 3}
	// fmt.Println("slice 1", s1)
	// s2 := []int{4, 5, 6, 7}
	// fmt.Println("slice 2", s2)
	// s1 = append(s1, s2...) //here ... is unpack operator which unpacks a slice
	// fmt.Println(s1)
	// fmt.Println("slice 1 after appending", s1)

	////copy of slice
	////func copy(dst []Type, src []Type) int --- returns length of new slice
	// s1 := []int{1, 2, 3}
	// s2 := []int{4, 5, 6, 4}
	// fmt.Println("slices before copying")
	// fmt.Println(s1)
	// fmt.Println(s2)
	// copy(s1, s2)
	// fmt.Println("slice1 after copying")
	// fmt.Println(s1)

	////NOTE:
	////nil slice vs empty slice
	////There is a difference between nil slice and an empty slice. nil slice is a slice with missing array reference
	////and empty slice is a slice with empty array reference or when the array is empty.

	////creating slice using make, it creates an empty slice not nil slice
	// s1 := make([]int, 2, 3)
	// fmt.Println(s1)

	////slices are passed by value to functions but since it points to the same array it feels like it is passed by reference

	////deleting an element from slice
	// s1 := []int{1, 2, 3, 4}
	// s1 = append(s1[:1], s1[2:]...)
	// fmt.Println(s1)

	/*
		slice of a large array can be bad for memory because garbage collector
		will not delete an array since a slice is pointing to it to avoid this can
		we can make a copy of slice upto needed data
	*/
	// s1 := []int{1, 2, 3, 4, 5} //can be many more this line will create an anonymous array
	// c1 := make([]int, len(s1))
	// copy(c1, s1)
	// fmt.Println((c1)) //now we can happily use this c1 slice and we dont need to worry about the large array

	////solving slice capacity increase mystery with append
	// arr1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr1)
	// s1 := arr1[1:3]
	// fmt.Println(s1, cap(s1))
	// s1 = append(s1, 0, 4, 5, 6, 7, 8, 9)

	// fmt.Println(s1, cap(s1))
	// fmt.Println(arr1)
}
