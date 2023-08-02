package lexer

import "arthur/token"

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

func (l *Lexer) NextToken() token.Token {
	var item token.Token

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
	}

	l.readChar()
	return item
}
