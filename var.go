package main

import f "fmt"

var (
	a int       = 3
	b string    = "Go!"
	c bool      = true
	d complex64 = -1 + 3i
)

func main() { // initial execution
	f.Println("Data Type: %T Value: %v\n", a, a)
	f.Println("Data Type: %T Value: %v\n", b, b)
	f.Println("Data Type: %T Value: %v\n", c, c)
	f.Println("Data Type: %T Value: %v\n", d, d)
}
