// Variadic function is a function that has only 1 parameter but have variable number of arguments
// Example: Print(), Println(), Printf() --> build-in finction from "fmt" package
// Please note that there is no variadic function in C but can can find this function in Java, JavaScript, Python, and etc.

package main

import f "fmt"

// create our own function to display data
func display(s ...string) {
	f.Println(s)
}

// main function to call display() function
func main() {
	display()                                  // [] --> No argument
	display("brown")                           // [brown] --> 1 argument
	display("black", "brown", "blue", "green") // [black brown blue green] --> 4 argument
}
