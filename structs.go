package main

import (
"fmt"
)

func main() {
//Struct example
type Person struct {
	name string
	addr string
	phone string
}
// Each property is a field, p1 contains values for all fields
var p1 Person 
var x string = "b st."

//Accessing Struct fields - Use dot notation
p1.name = "joe" // assign string "joe" to name field of p1
fmt.Println(p1.name)
p1.addr = x // assign x to be the address of p1
fmt.Println(p1.addr)

//Initializing Structs, can use new() - Initializes fields to zero
p2 := new(Person) // Makes a new Person thats empty, name address and phone will be empty strings
fmt.Println(p2) 
// Can initialize using a struct literal
p3 := Person{
	name: "john", 
	addr: "a st.", 
	phone: "123"}
fmt.Println(p3)
}