package main

import (
"fmt"
)

func main() {
//slices
slice := []string{"a", "b", "c"}

//accessing Length and Capacity
a1 := [3]string{"a", "b", "c"}
//length 1 because it only counts 0, as it finishes on 1. Capacity 3 because it has 3 elements
slice = a1[0:1]
fmt.Println(len(slice), cap(slice))

//Variable Slices
//Make - create a slice (and array)
//2 argument version - specify type (int,string) and length/capacity
sli := make([]int, 10)
fmt.Println(sli)
//3 argument version - specify type (int, string) and length and capacity separately
sli = make([]int, 10, 15)
fmt.Println(sli)

//Append - size of a slice can be increased, adds elements ot the end of a slice - inserts into the underlying array - increases the size of the array if necessary
sli = append(sli, 100)
fmt.Println(sli)
}