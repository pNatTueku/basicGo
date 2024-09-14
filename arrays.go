package main

func main() {
	var myArr = [3]string{"1st", "2nd", "3rd"}

	myArrCopy := myArr

	myArr[2] = "Another"

	myArr[2]
	myArrCopy[2]
}
