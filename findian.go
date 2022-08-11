package main

import (
"fmt"
"strings"
)

func main() {
// Defining input as string
var userInput string
// Prompting user to enter a string
fmt.Println("Enter a string:")
// Scanning user input
fmt.Scan(&userInput)
// Defining user input as lower case to make sure upper case letters also trigger the "Found!" message
userInput = strings.ToLower(userInput)
// If else statement defining that if the string starts with "i", contains "a" and ends with "n" it prints "Found!", otherwise, "Not Found!" is printed
if strings.HasPrefix(userInput, "i") && strings.Contains(userInput, "a") && strings.HasSuffix(userInput, "n") {
	fmt.Println("Found!")
} else {
	fmt.Println("Not Found!")
}
}