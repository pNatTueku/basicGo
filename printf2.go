// Printf() functon from fmt package can print as string template

package main

import f "fmt"

func main() {
	num := 24
	f.Printf("The decimal number is %d", num) // %s means that we are going to pass a decimal number to %d like here we pass num value to %s
}
