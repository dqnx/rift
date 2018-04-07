package main

import "gitlab.com/rauko1753/gorl"

// TermDrawText calls TermDraw multiple times to draw multiple characters.
func TermDrawText(x int, y int, s string) {
	for i, r := range s {
		gorl.TermDraw(x+i, y, gorl.Char(rune(r)))
	}
}
