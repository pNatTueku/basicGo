// Let's see the block function

package main

import f "fmt"

func main() {
	num := 10
	{
		num := 100
		f.Println(num) // first return 100 which in the inner here
	}
	f.Println(num) // second return 10 which is the outer
}
