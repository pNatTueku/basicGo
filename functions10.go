// Variadic function is a function that has only 1 parameter but have variable number of arguments
// Example: Print(), Println(), Printf() --> build-in finction from "fmt" package
// Please note that there is no variadic function in C but can can find this function in Java, JavaScript, Python, and etc.
// If we have both normal argument and variadic argument, then variadic argument need to be located at the last one

package main

import f "fmt"

// create our own function to display data
func display(s string, value ...int) {
	f.Println(s, value)
}

// main function to call display() function
func main() {
	display("none")                 // none [] --> No argument
	display("brown", 1, 2, 3, 4, 5) // brown [1, 2, 3, 4, 5]
	display("positive", -1, 0, 1)   // positive [-1, 0, 1]
}
