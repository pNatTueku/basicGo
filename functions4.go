// naked return - can be used only in the block function

// import package and build-in functions
package main

import (
	"fmt"
	"strings"
)

// create our own function to split strings
func split(str string) (s1 string, s2 string) {
	arr := strings.Split(str, ",")
	s1, s2 = arr[0], arr[1]

	return // return s1 and s2
}

// call our function and show the result
func main() {
	fmt.Println(split("Hello,Natty"))
}
