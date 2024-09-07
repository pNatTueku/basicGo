/*
AND: &&
OR: ||
NOT: !
*/

package main

import f "fmt"

func main() {

	var a, b, c = 1, 3, 5
	f.Println(a < b && a > c)
	f.Println(a < b || a > c)
	f.Println(!(a < b))
}
