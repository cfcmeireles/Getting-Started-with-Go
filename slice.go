package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
//Defining input as a string
var input string
//Making an empty slice of integers of size length 3
sli := make([]int, 3)

//For loop prompting user to enter an integer
for {
fmt.Println("Please enter an integer or enter x to exit:")
fmt.Scan(&input)
//If the user enters the character "x", the program quits
if input == "x" {
	fmt.Println("Exiting program")
	break
}
//Converting input to userInput with int type
userInput, _ := strconv.Atoi(input)
//Appending slice with userInput
sli = append(sli, userInput)
//Sorting integers in order and printing slice
sort.Ints(sli[:])
fmt.Println(sli)
 }
}