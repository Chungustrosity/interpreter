package lexer

import (
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // current position in the input (points to current char)
	readPosition int  // next position in the input    (after current char)
	rune         rune // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	jumpSize := 1
	if l.readPosition >= len(l.input) {
		l.rune = 0
	} else {
		l.rune, jumpSize = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += jumpSize
}
