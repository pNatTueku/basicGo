package main

import (
	"fmt"
	"math"
	r "math/rand" // set r as "math/rand"
	"time"
) // We can import mutilple packages by put them in ()

func main() {
	fmt.Println("Time is", time.Now())          // get time now
	fmt.Println("Sqrt(9) is", math.Sqrt((9)))   // find Sqrt
	fmt.Println("Random number is", r.Intn(10)) // random a number between [0, 10)
	fmt.Println("Pi is", math.Pi)
}
