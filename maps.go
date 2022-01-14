package main

//map is like a dictionary with dynamic types

func main() {
	////declaring an nil map object
	////this line created a nil map
	// var m map[string]int
	////this will throw an error since this is nil map
	// m["num"] = 1

	////declaring an empty map object
	// m1 := make(map[string]int)
	// m1["numbe"] = 1 // this will work fine
	// fmt.Println(m1)

	////assigning value with declaration
	// m2 := map[string]int{
	// 	"number": 1,
	// }
	// fmt.Println((m2))

	////accessing map data
	// m3 := map[string]int{
	// 	"number": 1,
	// }
	// fmt.Println(m3["number"])
	////if value is not in map we can verify it using a tuple
	// value, isexists := m3["number"]
	// fmt.Println(value, isexists)
	// value1, isexists1 := m3["number1"]
	// fmt.Println(value1, isexists1)

	////length of map
	// fmt.Println(len(m3))

	////delete map object
	// delete(m3, "number")
	// fmt.Println(m3)

	////iterating maps
	// for key, value := range m3 {
	// 	fmt.Println(key, "===>", value)
	// }

	////map with other datatypes
	// m3 := map[bool]int{
	// 	true: 1,
	// }
	// fmt.Println(m3[true])

	////NOTE: maps cannot have repeated key values
	////if we assign a map to other map it will simply refer to the internal datastructure just like a slice
	////to make a deep copy of map we need to iterate it

	////this creates a deepcopy of map
	// newmap := make(map[bool]int)
	// for key, value := range m3 {
	// 	newmap[key] = value
	// }
	// fmt.Println(newmap)

}
