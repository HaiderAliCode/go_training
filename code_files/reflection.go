package main

import (
	"fmt"
	"reflect"
)

//build around 3 concepts:
//types, kinds, values

type employee struct {
	name string
	age  int
}

func main() {
	// var s int
	// s = 0
	// fmt.Println(reflect.TypeOf(s))
	// fmt.Println(reflect.TypeOf(s).Name())
	// fmt.Println(reflect.Kind(s))

	emb := employee{
		name: "Jadoo",
		age:  15,
	}

	// return name of new type
	fmt.Println("Type of struct", reflect.TypeOf(emb))
	//return all values
	fmt.Println("value of struct", reflect.ValueOf(emb))

	//return struct
	fmt.Println("Value of .kind", reflect.ValueOf(emb).Kind())

	if reflect.ValueOf(emb).Kind() == reflect.Struct {
		v := reflect.ValueOf(emb)
		for i := 0; i < v.NumField(); i++ {
			//printting all fields
			fmt.Println(v.Field(i))
		}
	}
	// fmt.Println(reflect.TypeOf(emb))

	b := 15
	//can help extract value in int or string
	fmt.Println(reflect.ValueOf(b).Int())
	fmt.Println(reflect.ValueOf(b).String())

	// Clear is better than clever. Reflection is never clear.
	// Reflection is a very powerful and advanced concept in Go and it should be used with care.

	//can compare multiple structs with reflect.DeepEqual(map_1, map_2)

}
