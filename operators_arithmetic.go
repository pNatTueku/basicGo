package main

import f "fmt"

func main() {
	var x, y = 10, 20
	f.Printf("x + y = %d\n", x+y)   // additive
	f.Printf("x - y = %d\n", x-y)   // abstract
	f.Printf("x * y = %d\n", x*y)   // multiplication
	f.Printf("x / y = %d\n", x/y)   // Deviation
	f.Printf("x mod y = %d\n", x%y) // modulation
	x++
	f.Printf("x++ = %d\n", x)
	y--
	f.Printf("y-- = %d\n", y)
}
