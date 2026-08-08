package lexer

import (
	"interpreter/token"
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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.rune {
	case '=':
		tok = newToken(token.ASSIGN, string(l.rune))
	case '+':
		tok = newToken(token.PLUS, string(l.rune))
	case ',':
		tok = newToken(token.COMMA, string(l.rune))
	case ';':
		tok = newToken(token.SEMICOLON, string(l.rune))
	case '(':
		tok = newToken(token.LPAREN, string(l.rune))
	case ')':
		tok = newToken(token.RPAREN, string(l.rune))
	case '{':
		tok = newToken(token.LBRACE, string(l.rune))
	case '}':
		tok = newToken(token.RBRACE, string(l.rune))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.rune) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupKeyword(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, string(l.rune))
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, tokenLiteral string) token.Token {
	return token.Token{Type: tokenType, Literal: tokenLiteral}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.rune) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// TODO: extend to correctly handle non-ASCII, unicode runes.
func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.rune == ' ' || l.rune == '\t' || l.rune == '\n' || l.rune == '\r' {
		l.readChar()
	}
}
