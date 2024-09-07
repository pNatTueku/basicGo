/*
Postfix: () [] <- . ++ -- 				Left to Right
Unary: + - ! ~ ++ -- (type)* & sizeof 	Right to Left
Additive: + - 		Left to Right
Multiplicative: * / % 		Left to Right
Shift: << >>				Left to Right
Relational: < <= > >=		Left to Right
Equality: == !=				Left to Right
Bitwise AND: &				Left to Right
Bitwise XOR: ^				Left to Right
Bitwise OR: |				Left to Right
Bitwise XOR: ^				Left to Right
Logical AND: &&				Left to Right
Logical OR: ||				Left to Right
Assignment: = += -= *= /= %= >>= <<= &= ^= |= Left to Right
Comma: , Left to Right
*/

package main

import f "fmt"

func main() {
	n := 5*2 + 10/2 - 1
	f.Println(n)
}
