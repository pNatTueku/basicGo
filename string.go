package main

import f "fmt"

func main() {
	f.Println(len("Hello Natty!\t\n")) // use  len() to count the length, here we also count a space and escape character "\t", "\n"
	f.Print("Go" + "lang" + "!")       // we can use "+" to concatenate strings
}
