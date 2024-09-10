package main

import f "fmt"

func main() {
	hi() // function hi() here can see function hi() from below script BTW in C it cannot see
}

func hi() {
	f.Println("Hi hi!")
}
