package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	userData := map[string]string{}


	// 1) get the user's name
	fmt.Print("Please Enter Your Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	userData["name"] = name

	// 1.1) get the user's address
	fmt.Print("Please Enter Your Current Address: ")
	_, _err := fmt.Scan(&address)
	if _err != nil {
		fmt.Println(_err)
	}
	userData["address"] = address

	// 2)
	// marshall final data:
	dataAsJson, __err := json.Marshal(userData)
	if __err != nil {
		fmt.Println(__err)
	}

	// 3)
	// represent final data	:
	fmt.Println()
	fmt.Println("========= Your data represented in json string =========")
	fmt.Println(string(dataAsJson))
}
