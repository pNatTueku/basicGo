package main

import "fmt"

func main() { // main() is an initial function to run program which cannot be changed to other name. BTW there is no return value from this main()
	fmt.Println("Hello GO's World")
}

func Main() { // Go is case sensitive - main() is different from Main()
	fmt.Println(("Bye"))
}
