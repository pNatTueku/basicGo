// rune aka int32 - it is a unicode code point
package main

import f "fmt"

var (
	a rune = 3
)

func main() { // initial execution
	f.Printf("Data Type: %T Value: %v\n", a, a) // printf formats specifier and writes the standard output
}
