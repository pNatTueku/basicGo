package main

import f "fmt"

func main() {
	var x, y = 3, 5
	x = y
	f.Println("x = y, x =", x)
	x = 3
	x += y
	f.Println("x += y, x =", x)
	x = 3
	x -= y
	f.Println("x -= y, x =", x)
	x = 3
	x *= y
	f.Println("x *= y, x =", x)
	x = 3
	x /= y
	f.Println("x /= y, x =", x)
	x = 3
	x %= y
	f.Println("x %= y, x =", x)
}
