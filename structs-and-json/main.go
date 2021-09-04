package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name  string
	Age   int
	Hired bool
	time  time.Time
	Boss  *Employee
}

func main() {

	company := make(map[string]*Employee)

	company["e2"] = &Employee{
		Name:  "Salih",
		Age:   22,
		Hired: false,
		time:  time.Now(),
		Boss:  company["e1"],
	}

	company["e1"] = &Employee{
		Name:  "Salih",
		Age:   22,
		Hired: false,
		time:  time.Now(),
	}

	fmt.Printf("%T %+[1]v\n", company["e1"])
	fmt.Printf("%T %+[1]v\n", company["e2"])

	fmt.Printf("%v\n", company)

	/*
		another use case for struct
		add value directly after initialization
	*/
	t := struct {
		name string
		age  int
	}{
		name: "Mohamed",
		age:  50,
	}

	fmt.Println(t)

	/*
		anonimouse structs are compatible is they follows the rules.
	*/

	/*
		structs are SINGLETONE
	*/

	s1 := struct {
		Z int `json:"z"` // `json:"XXXX"`  is called struct's tag
	}{1}

	s2 := struct {
		Z int `json:"z"`
	}{5}

	s1 = s2

	fmt.Println(s1) // {5}

	type Response struct {
		Value []string `json:"value"` // Value not value to be exported, It Must To Be UPPER case
		Num   int      `json:"num,omitempty"`
	}

	structWithTags := Response{Value: []string{"go", "to", "kitchen"}, Num: 1}
	parseStructWithTags, _ := json.Marshal(structWithTags) // encoding structs in json type

	// new instance of the struct, to save the data on it after pardsing.
	var originalWithNoParsing Response
	_ = json.Unmarshal(parseStructWithTags, &originalWithNoParsing) // dencoding json to struct result

	fmt.Println("structWithTags", structWithTags)
	fmt.Println("parseStructWithTags", string(parseStructWithTags))
	fmt.Println("originalWithNoParsing", originalWithNoParsing)

}
