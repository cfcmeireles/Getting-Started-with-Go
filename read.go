package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {

var userInput string

type FullName struct {
	fname string
	lname string
}
// Asking for file name
fmt.Println("Please enter file name:")
fmt.Scan(&userInput)

// Opening file
readFile, err := os.Open("File.txt")
if err != nil {
 fmt.Println("File not found")
}

// Slice containing struct
namesList := make([]FullName, 0)

// Scanning file
fileScanner := bufio.NewScanner(readFile)
for fileScanner.Scan() {

line := fileScanner.Text()
res := strings.Split(line, " ")

var name FullName
name.fname = res[0]
name.lname = res[1]
namesList = append(namesList, name)
}
// Iterating through namesList slice and printing first and last names found in each struct
for i := 0; i < len(namesList); i++ {
	fmt.Println("First name:", namesList[i].fname)
	fmt.Println("Last name:", namesList[i].lname)
	fmt.Println(" ")
}
}