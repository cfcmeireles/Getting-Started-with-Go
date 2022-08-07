package main

import (
"fmt"
)

func main() {
//arrays
//array of 5 integers
var x [5]int
//defining particular element "zero" and setting it to equal to 2
x[0] = 2
fmt.Println(x[0])

//array literal - pre-defined with values
t := [5]int{1, 2, 3, 4, 5}
fmt.Println(t) //prints array
// [...] will infer the size from number of initializers, in this case 4
w := [...]int{1, 2, 3, 4}
fmt.Println(w)
fmt.Println(len(w)) //prints length of array 4

//iterating through arrays
o := [3]int{1, 2, 3}
//using for loop
for index, element := range o {
	fmt.Println(index, element) //range sets the scope of iteration up to the length of the array -> prints the index of the array (0, 1, 2) and the element of the array (1, 2, 3)
}
for _, element := range o {
fmt.Println(element) // use blank identifier _ when we don't need index, prints only the elements of the array
}
}