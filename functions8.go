// Variadic function is a function that has only 1 parameter but have variable number of arguments
// Example: Print(), Println(), Printf() --> build-in finction from "fmt" package
// Please note that there is no variadic function in C but can can find this function in Java, JavaScript, Python, and etc.

package main

import f "fmt"

// create our own function to display data
func display(s ...string) {
	f.Println(s[0]) // first string
	f.Println(s[2]) // third string
}

// main function to call display() function
func main() {
	display("black", "brown", "blue", "green")
}
