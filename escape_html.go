// HTML's string has <, >, &, ', ", and etc. If we would like to encode we can use an EscapeString() function from html package

package main

import (
	"fmt"
	"html"
)

func main() {
	const x = "<div>Golang&Java</div>"
	fmt.Println(html.EscapeString(x)) // --> We will get &lt;div&gt;Golang&amp;Java&lt;/div&gt;
}
