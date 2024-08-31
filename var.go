package main

import f "fmt"

var (
	a int       = 3
	b string    = "Go!"
	c bool      = true
	d complex64 = -1 + 3i        // complex #1
	e complex64 = complex(-1, 3) // complex #2
)

func main() { // initial execution
	f.Printf("Data Type: %T Value: %v\n", a, a) // printf formats specifier and writes the standard output
	f.Printf("Data Type: %T Value: %v\n", b, b)
	f.Printf("Data Type: %T Value: %v\n", c, c)
	f.Printf("Data Type: %T Value: %v\n", d, d)
	f.Printf("Data Type: %T Value: %v\n", e, e)
}
