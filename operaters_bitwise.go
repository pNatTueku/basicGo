/*
AND between bit: &
OR between bit: |
XOR between bit: ^
Zero fill left shift: <<
Signed right shift: >>
*/

package main

import f "fmt"

func main() {
	var x uint = 10 // 0000 1010
	var y uint = 55 // 0011 0111
	var z uint

	z = x & y
	f.Println("x & y = ", z) // 0000 1010 &  0011 0111 = 0000 0010
	z = x | y
	f.Println("x | y = ", z) // 0000 1010 |  0011 0111 = 0011 1111
	z = x ^ y
	f.Println("x ^ y = ", z) // 0000 1010 ^  0011 0111 = 0011 1101
	z = x << 1
	f.Println("x << 1 = ", z) // 0000 1010 << 1 = 0001 0100
	z = x >> 1
	f.Println("x >> 1 = ", z) // 0000 1010 << 1 = 0000 0101
}
