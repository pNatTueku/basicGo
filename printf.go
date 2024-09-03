// Printf() functon from fmt package can print as string template

package main

import f "fmt"

func main() {
	msg := "Hello Natty!"
	f.Printf("The string is %s", msg) // %s means that we are going to pass a string to %s like here we pass msg value to %s
}
