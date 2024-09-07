// Get the user input from console,require  enter, and then it works

package main

import (
	f "fmt"
)

func main() {
	var integer_value int
	var string_value string
	var float_value float64
	var boolean_value bool

	f.Scanf("%d", &integer_value)
	f.Scanf("%s", &string_value)
	f.Scanf("%g", &float_value)
	f.Scanf("%t", &boolean_value)
	f.Printf("%d %s %g %t", integer_value, string_value, float_value, boolean_value)
}
