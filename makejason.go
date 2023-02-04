/*Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.*/


package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	var personName string
	var adrs string

	fmt.Println("Enter a name")

	fmt.Scan(&personName)

	fmt.Println("Enter the address ")

	fmt.Scan(&adrs) 

	var userTest = map[string]string{"fname": personName, "address": adrs} 

	userTestJSON, err := json.Marshal(userTest) 

	if err != nil {
		fmt.Println("error whie marshalling userTest map")
	}

	fmt.Println("Jason: ", string(userTestJSON)) 

}
