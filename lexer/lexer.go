package lexer

import (
	"arthur/token"
)

// TODO: change lexer to support unicode
type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// returns true if the character being read is a letter
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_' ||
		char == '!' ||
		char == '\\'
}

// completely reads keywords and user-defined identifiers
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// returns true if the character being read is a digit
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

// completely reads a number while it doesn't find any digit anymore
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// doesn't skip newline since it replaces semicolons
func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

// returns the correct token
func (l *Lexer) NextToken() token.Token {
	var item token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '=':
		item = newToken(token.ASSIGN, l.char)
	case ';':
		item = newToken(token.SEMICOLON, l.char)
	case '(':
		item = newToken(token.LPAREN, l.char)
	case ')':
		item = newToken(token.RPAREN, l.char)
	case ',':
		item = newToken(token.COMMA, l.char)
	case '+':
		item = newToken(token.PLUS, l.char)
	case '{':
		item = newToken(token.LBRACE, l.char)
	case '}':
		item = newToken(token.RBRACE, l.char)
	case '\n':
		item = newToken(token.NEWLINE, l.char)
	case 0:
		item.Literal = ""
		item.Type = token.EOF
	default:
		if isLetter(l.char) {
			item.Literal = l.readIdentifier()
			item.Type = token.LookupIdent(item.Literal)
			return item
		} else if isDigit(l.char) {
			item.Literal = l.readNumber()
			item.Type = token.INT
			return item
		}
		item = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return item
}
