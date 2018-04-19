package main

import "gitlab.com/rauko1753/gorl"

// TODO: Paramaterize these functions to accept keymaps for different
// scenes.

// KeyToVector converts a key press to a unit vector direction.
func KeyToVector(k gorl.Key) (bool, gorl.Offset) {
	var o gorl.Offset
	switch k {
	case 'u':
		o = gorl.Offset{-1, 1}
	case 'i':
		o = gorl.Offset{0, 1}
	case 'o':
		o = gorl.Offset{1, 1}
	case 'j':
		o = gorl.Offset{-1, 0}
	case 'k':
		o = gorl.Offset{0, 0}
	case 'l':
		o = gorl.Offset{1, 0}
	case 'm':
		o = gorl.Offset{-1, -1}
	case ',':
		o = gorl.Offset{0, -1}
	case '.':
		o = gorl.Offset{1, -1}
	default:
		return false, o
	}

	return true, o
}
