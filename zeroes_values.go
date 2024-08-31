/*
Zeroes values are the default value
- numeric: 0
- bool: false
- string: ""
- pointer: nil
- slice: nil
- map: nil
- channel: nil
- function: nil
- interface: nil
*/

package main

import f "fmt"

func main() {
	var a int
	var b bool
	var c complex64
	var d string
	var e float64

	f.Printf("%v, %v, %v, %q, %v", a, b, c, d, e)
	f.Println()
	f.Printf("%v, %v, %v, %v, %v", a, b, c, d, e)
}
