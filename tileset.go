package main

import (
	"bytes"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// StringCP437 encodes a string as a buffer of ints in CodePage437.
func StringCP437(s string) []byte {
	var b bytes.Buffer
	encoder := transform.NewWriter(&b, charmap.CodePage437.NewEncoder())
	encoder.Write([]byte(s))
	encoder.Close()

	return b.Bytes()
}

func RuneCP437(r rune) byte {
	c, _ := charmap.CodePage437.EncodeRune(r)
	return c
}
