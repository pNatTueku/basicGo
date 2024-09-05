// Raw string literal is used to write string by using backticks ``. Those string will not be complied (escape character) and also can start the new line in that ``

package main

import f "fmt"

func main() {
	msg := `<div>
	line1
	</div>`
	f.Println(msg)
}
