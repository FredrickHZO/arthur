package lexer

import (
	"arthur/token"
	"unicode"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
		return
	}
	l.char = rune(l.input[l.readPosition])
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, char rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// reads keywords and user-defined identifiers
func (l *Lexer) lexIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// reads a number until there are no digit left
func (l *Lexer) lexNumber() string {
	startPos := l.position
	for unicode.IsDigit(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// doesn't skip newline since it replaces semicolons
func (l *Lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

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
		// might break for _ ! \
		if unicode.IsLetter(l.char) {
			item.Literal = l.lexIdentifier()
			item.Type = token.LookupIdent(item.Literal)
			return item
		}
		if unicode.IsDigit(l.char) {
			item.Literal = l.lexNumber()
			item.Type = token.INT
			return item
		}
		item = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return item
}
