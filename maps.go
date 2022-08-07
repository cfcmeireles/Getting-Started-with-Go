package main

import (
"fmt"
)

func main() {
//Implementation of a hash table
var idMap map[string]int
//Use make() to create a map -> in this example, string is key type, int is value type
idMap = make(map[string]int)
//Define a map literal -> joe is the key, 123 is the value.
idMap = map[string]int {
	"joe": 123,
    "jane": 567 }
fmt.Println(idMap) //Prints the key and values inside the map

//Accessing Maps
fmt.Println(idMap["joe"]) // Prints the value associated with joe, in this case 123

//Adding a key/value pair into the map
idMap["jane"] = 456 // It can add a new key/value into the map OR replace the existing one
fmt.Println(idMap["jane"]) // Prints the new/replaced value associated with jane, in this case 456
fmt.Println(idMap) // Now prints the map with the new added OR replaced key "jane" and its new value 456

//Deleting a key/value pair of a map
delete(idMap, "joe")
fmt.Println(idMap) // Prints the map without the deleted key/value of "joe"

//Map functions
//Two-value assignment
id, p := idMap["jane"] // id is value, p is presence of key (boolean)
fmt.Println(id, p) // prints value "456" and "true" for presence of key - if I used "joe" instead, it would print "0" and "false" because it was deleted and doesn't exist anymore

// len() returns number values
fmt.Println(len(idMap))

//Iterating through a map - Use a for loop with the range keyword. Two-value assignment with range
for key, val := range idMap {
	fmt.Println(key, val) // Prints the key and value inside of the idMap
}
}