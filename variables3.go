package main

import f "fmt"

func main() {
	i := 6                       // short declaration - no need var i int = 6
	a, b, c := 7, "Natty", false // var a, b, c = 7, "Natty", false
	f.Println(a, b, c, i)
}
