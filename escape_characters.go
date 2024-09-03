/*
Escape character is the special character started with "\"
- \a : Alert or Bell
- \b : Backspace
- \\ : Backslash
- \t : Horizontal tab
- \n : Line feed or new line
- \f : Form feed
- \r : Carriage return
- \v : Vertical tab
- \' : Single quote which will be used only in Rune
- \" : Double quote which will be used only in String

Here are the escape character for encoding
- \x : follows with 2 digits of Hexadenary number
- \u : follows with 4 digits of Hexadenary number
- \U : follows with 8 digits of Hexadenary number
- \  : follows with 3 digits of Octanary number
*/

package main

import f "fmt"

func main() {
	f.Println("\x4f\x6b\x21")
	f.Println("\117\153\041")
	f.Println("caf\u00e9")
	f.Println("I\tlove\tU")
}
