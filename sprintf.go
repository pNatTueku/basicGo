// In fmt package, we can define string template by using Sprintf() functiom

package main

import f "fmt"

func main() {
	s := f.Sprintf("X = %d", 123)
	f.Println(s)
}
