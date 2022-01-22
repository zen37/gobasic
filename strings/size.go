package main

/* A rune is an alias to the int32 data type (analogous to byte as an alias for uint8). It represents a Unicode code point.
A Unicode code point or code position is a numerical value that is usually used to represent a Unicode character.
The int32 is big enough to represent the current volume of 140,000 unicode characters.
ASCII defines 128 characters, identified by the code points 0–127.
Unicode, a superset of ASCII, defines the codespace of 1,114,112 code points. */

import (
	"fmt"
	"unicode/utf8"
)

var p = fmt.Println

func main() {
	ch := "世界"
	en := "world"
	bg := "свят"
	jp := "世界"
	ab := "العالمية" //arabic

	p("lengh of", ch, "is", len(ch))
	p("lengh of", en, "is", len(en))
	p("lengh of", bg, "is", len(bg))
	p("lengh of", jp, "is", len(jp))
	p("lengh of", ab, "is", len(ab))
	p()
	p("rune lengh of", ch, "is", utf8.RuneCountInString(ch))
	p("rune lengh of", en, "is", utf8.RuneCountInString(en))
	p("rune lengh of", bg, "is", utf8.RuneCountInString(bg))
	p("rune lengh of", jp, "is", utf8.RuneCountInString(jp))
	p("rune lengh of", ab, "is", utf8.RuneCountInString(ab))
}
