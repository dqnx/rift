package main

import (
	"bytes"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// StringCP437 encodes a string as a buffer of ints in CodePage437.
func StringCP437(s string) bytes.Buffer {
	var b bytes.Buffer
	encoder := transform.NewWriter(&b, charmap.CodePage437.NewEncoder())
	encoder.Write([]byte(s))
	encoder.Close()

	return b
}
