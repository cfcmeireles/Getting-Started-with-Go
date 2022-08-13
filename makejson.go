package main

import (
"fmt"
"encoding/json"
)

func main() {
//Declaring name and address input as variables of string type
var nameInput string
var addrInput string
//Declaring idMap as a variable with the key and value of string type
var idMap map[string]string
//Prompting user to enter name and address and scanning nameInput and addrInput
fmt.Println("Please enter your name:")
fmt.Scan(&nameInput)
fmt.Println("Please enter your address:")
fmt.Scan(&addrInput)
//Creating idMap
idMap = make(map[string]string)
//Defining "name" and "address" as keys in the map with nameInput and addrInput as values
idMap = map[string]string {
	"name": nameInput,
    "address": addrInput}
//Creating a JSON object from the idMap
jsonObj, _ := json.Marshal(idMap)
//Printing the JSON object
fmt.Println(string(jsonObj))
}