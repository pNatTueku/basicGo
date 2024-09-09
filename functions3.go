// get input and return multiple result

package main

import f "fmt"

// create our worn function
func swap(x, y int) (int, int) {
	return y, x
}

// call our function and show the results
func main() {
	x, y := 10, 20
	x, y = swap(x, y)
	f.Println(x, y) // We will see the returns value as 20 10

}
