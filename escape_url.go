// HTML's string has a space, /, ?, and etc. If we would like to encode we can use a PathEscape() function from net/url package

package main

import (
	"fmt"
	"net/url"
)

func main() {
	const x = "<div>Golang&Java</div>"
	fmt.Println(url.PathEscape(x)) // --> We will get %3Cdiv%3EGolang&Java%3C%2Fdiv%3E
}
