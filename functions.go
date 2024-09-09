package main

import f "fmt"

// require parameter and data type, and data type to be returned
func area(a int, b int) int {
	return a * b
}

// call our function by set input values and print result here
func main() {
	result := area(4, 5)         // input a and b then call area() function and store the value in result
	f.Println("Area = ", result) // print result
}
