package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

//please remember to capitalize first letter of struct properties
type student struct {
	Name           string
	Ismale         bool
	Age            int
	Heightinmeters float64
}

//maps to json
type student2 map[string]interface{}

//nested json
type profile struct {
	Subjectname string
}

type student3 struct {
	Name    string
	Profile profile
}

//json pointer interface value
type ProfileI interface {
	Check() string
}

type Profile struct {
	Name string
	Age  int
}

func (p Profile) Check() string {
	return "checked"
}

type student4 struct {
	Name    string
	Profile ProfileI
}

//struct tags
type student5 struct {
	Name string `json:"fname"`
	//omitempty removes label if the value inside it is empty
	LName string `json:"lname,omitempty"`
}

type student6 struct {
	Name string `json:"fname"`
}

func main() {
	john := student{
		Name:           "Haider",
		Ismale:         false,
		Age:            25,
		Heightinmeters: 1.6,
	}

	fmt.Println(john)
	json_str, _ := json.Marshal(john)
	fmt.Println(string(json_str))

	//map to json
	johna := student2{
		"name": "Haider",
		"age":  2,
	}

	fmt.Println(johna)
	json_stra, _ := json.Marshal(johna)
	fmt.Println(string(json_stra))

	//nested json object
	std3 := student3{
		Name:    "Haider ali",
		Profile: profile{Subjectname: "golang"},
	}
	fmt.Println(std3)
	json_strstd3, _ := json.MarshalIndent(std3, "", "  ")
	fmt.Println(string(json_strstd3))

	// Note: same like structs anonymous struct transfer data into parent

	//pointers with json and structs
	emp5 := student4{
		Name:    "Haider Ali",
		Profile: Profile{Name: "Haider", Age: 17},
	}
	fmt.Println(emp5)

	emp6 := student5{
		Name: "Haider Ali",
		// LName: "M. ",
	}
	json_stremp6, _ := json.Marshal(emp6)
	fmt.Println(string(json_stremp6))

	//////////Decoding json//////////
	//checking if json is valid

	data_json := []byte(`
	{
		"Name": "Haider"
	}
	`)
	//return true if valid else false
	fmt.Println(json.Valid(data_json))

	var cp Profile
	json.Unmarshal(data_json, &cp)
	//marshalling back
	jso_str, _ := json.Marshal(cp)
	fmt.Println(string(jso_str))

	//Notes:Hence, it’s better to designate a number field as float32/64 in scenarios where a field can either be an integer or a floating-point number.
	//some important unmarshalling errors
	// 	The fields that are unexported in the struct or missing from the JSON are not unmarshalled. If a field value in the JSON is null and its corresponding field type’s zero-value is nil (like interface, map, pointer, or slice) then the value is replaced with nil, else that field is ignored for unmarshalling and retains its original value.
	// If Unmarshal encounters an array type and array values in the JSON are more than the array can hold, then extra values are discarded. If array values in JSON are less than the length of the array, then the remaining array elements are set to their zero-values. The array type should be compatible with the values in the JSON.
	// If Unmarshal encounters a slice type, then the slice in the struct is set to 0 length and elements from the JSON array are appended one at a time. If the JSON contains an empty array, then Unmarshal replaces the slice in the struct with an empty slice. The slice type should be compatible with the values in the JSON.
	// If Unmarshal encounters a map type and the map’s value in the struct is nil, then a new map is created and object values in the JSON are appended. If the map value is non-nil, then the original value of the map is reused and new entries are appended. The map type should be compatible with the values in the JSON.
	// If Unmarshal encounters a pointer field and the value of that field in the JSON is null then that field is set to nil pointer value. If the field in the JSON is not null then new memory is allocated for the pointer in case the pointer is nil or the old value of the pointer is reused.

	// if the struct have nested structs defined than the json should have fields in root/ open else it wont unmarshal

	var std6 student6
	data := []byte(`
	{
		"fname":"Haider Ali"
	}
	`)
	json.Unmarshal(data, &std6)
	json_str_std6, _ := json.Marshal(std6)
	fmt.Println(string(json_str_std6))
	fmt.Println(std6)

	//or we can parse using map
	type dataparser map[string]interface{}
	var dataparsed dataparser

	json.Unmarshal(data, &dataparsed)
	json_str_std6_map, _ := json.Marshal(dataparsed)
	fmt.Println(string(json_str_std6_map))
	fmt.Println(dataparsed)

	//or simply a data container
	var datacon interface{}
	json.Unmarshal(data, &datacon)
	fmt.Println(datacon)
	//throws error
	// fmt.Println(datacon["fname"])
	//fixing error
	dataconfixed := datacon.(map[string]interface{})
	fmt.Println(dataconfixed["fname"])

	//buffer encoder
	buff := new(bytes.Buffer)
	bufEncorder := json.NewEncoder(buff)
	bufEncorder.Encode(student6{Name: "Haider Ali"})
	bufEncorder.Encode(student6{Name: "Tabish Amin"})
	fmt.Println(buff)

	//buffer decoder
	jsonstream := strings.NewReader(`
	{"fname": "Haider Ali"}
	{"fname": "Tabish Amin"}
	`)

	decoder := json.NewDecoder(jsonstream)

	var hdata, tdata student6

	decoder.Decode(&hdata)
	decoder.Decode(&tdata)
	fmt.Println(hdata)
	fmt.Println(tdata)

}
