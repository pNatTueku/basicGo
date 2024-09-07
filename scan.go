package main

import f "fmt"

func main() {
	var num int

	f.Printf("Please enter an integer:")
	f.Scan(&num)
	f.Printf("A number is %d", num)
}
