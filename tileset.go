package main

import (
	"bytes"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// CP437 encodes a string as a buffer of ints in CodePage437.
func CP437(s string) bytes.Buffer {
	var b bytes.Buffer
	encoder := transform.NewWriter(&b, charmap.CodePage437.NewEncoder())
	encoder.Write([]byte(s))
	encoder.Close()

	return b
}

// Tile is a "char" enum.
type Tile int

const (
	Equal3 Tile = iota
	PlusMinus
	GreatEquals
	LessEquals
	ReverseHook
	Hook
	Divide
	Tilde2
	Degree
	Dot
	SmallDash
	Root
	SmallN
	Exp2
	BigDot
	Blank
	Alpha
	Beta
	Gamma
	Pi
	UpperSigma
	Sigma
	Mu
	Tau
	Phi
	Theta
	Omega
	delta
	Inf
	Ninf
	Euro
	Intersect
	W2up1h
	W1down2h
	W2down1h
	W2up1right
	W1up2right
	W1down2right
	W2down1right
	W2v1h
	W1v2h
	W1left1up
	W1down1right
	Fill100
	FillBot
	FillLeft
	FillRight
	FillTop
	W1up1right
	W1up1h
	W1down1h
	W1v1right
	W1h
	W1v1h
	W1v2right
	W2v1right
	W2up2right
	W2down2right
	W2up2h
	W2down2h
	W2v2right
	W2h
	W2v2h
	W1up2h
	Fill25
	Fill50
	Fill75
	W1v
	W1v1left
	W1v2left
	W2v1left
	W2down1left
	W1down2left
	W2v2left
	W2v
	W2down2left
	W2up2left
	W2up1left
	W1up2left
	W1down1left
	a_acute
	i_acute
	o_acute
	u_acute
	n_tilde
	N_tilde
	a_under
	o_under
	QuestionRev
	HyphenLeft
	HyphenRight
	Half
	Quarter
	ExclamationRev
	Angle2Left
	Angle2Right
	E_acute
	ae
	AE
	o_circ
	o_umlaut
	o_grave
	u_circ
	u_umlaut
	y_umlaut
	O_umlaut
	U_umlaut
	Cent
	Pound
	Yen
	Pt
	Func
	C_cedilla
	mu_umlaut
	e_acute
	a_circ
	a_umlaut
	a_grave
	a_invbreve
	c_cedilla
	e_circ
	e_umlaut
	e_grave
	i_umlaut
	i_circ
	i_grave
	A_umlaut
	A_invbreve
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	BraceLeft
	Pipe
	BraceRight
	Tilde
	Delta
	Backtick
	a
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	BracketLeft
	Slash
	BracketRight
	Caret
	Underscore
	At
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	Zero
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Colon
	Semicolon
	AngleLeft
	Equals
	AngleRight
	Question
	Blank2
	Exclamation
	Quote2
	Hash
	Dollar
	Percent
	Ampersand
	Apostrophe
	ParenLeft
	ParenRight
	Asterisk
	Plus
	Comma
	Dash
	Period
	Backslash
	TriRight
	TriLeft
	TriUpDown
	Exclamation2
	Paragraph
	Subsection
	FatDash
	TriBoxUp
	ArrowUp
	ArrowDown
	ArrowRight
	ArrowLeft
	SubDash
	ArrowLeftRight
	TriUp
	TriDown
	Blank3
	Person1
	Person2
	Heart
	Diamond
	Club
	Spade
	OvalFill
	OvalFillRev
	Circle
	CircleRev
	Male
	Female
	Music1
	Music2
	Star
)
