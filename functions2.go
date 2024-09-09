package main

import f "fmt"

// write our function to get inputs but no return, print here
func show_area(a int, b int) {
	result := a * b
	f.Println("Area = ", result)
}

// call function
func main() {
	show_area(4, 5)
}
