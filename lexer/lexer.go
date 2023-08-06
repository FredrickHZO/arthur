package lexer

import (
	"arthur/token"
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

func (l *Lexer) peek() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
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

// reads keywords and user-defined identifiers
func (l *Lexer) lexIdentifier() string {
	position := l.position
	for isLetter(l.char) || isDigit(l.char) {
		if l.position == len(l.input)-1 {
			l.position = l.readPosition
			break
		}
		l.readChar()
	}
	return l.input[position:l.position]
}

// reads a number until there are no digit left
func (l *Lexer) lexNumber() string {
	startPos := l.position
	for isDigit(l.char) {
		if l.position == len(l.input)-1 {
			l.position = l.readPosition
			break
		}
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
	// operator cases
	case '=':
		if l.peek() == '=' {
			item = newCompToken(token.EQ, "==")
			l.readChar()
		} else {
			item = newToken(token.ASSIGN, l.char)
		}
	case '+':
		if l.peek() == '+' {
			item = newCompToken(token.INCREMENET, "++")
			l.readChar()
		} else if l.peek() == '=' {
			item = newCompToken(token.PLUS_EQ, "+=")
			l.readChar()
		} else {
			item = newToken(token.PLUS, l.char)
		}
	case '-':
		if l.peek() == '-' {
			item = newCompToken(token.DECREMENT, "--")
			l.readChar()
		} else if l.peek() == '=' {
			item = newCompToken(token.MINUS_EQ, "-=")
			l.readChar()
		} else {
			item = newToken(token.MINUS, l.char)
		}
	case '!':
		if l.peek() == '=' {
			item = newCompToken(token.NOT_EQ, "!=")
			l.readChar()
		} else {
			item = newToken(token.BANG, l.char)
		}
	case '*':
		if l.peek() == '=' {
			item = newCompToken(token.ASTERISK_EQ, "*=")
			l.readChar()
		} else {
			item = newToken(token.ASTERISK, l.char)
		}
	case '/':
		if l.peek() == '=' {
			item = newCompToken(token.SLASH_EQ, "/=")
		} else {
			item = newToken(token.SLASH, l.char)
		}
	case '<':
		item = newToken(token.LT, l.char)
	case '>':
		item = newToken(token.RT, l.char)
	// delimiter cases
	case ',':
		item = newToken(token.COMMA, l.char)
	case ';':
		item = newToken(token.SEMICOLON, l.char)
	case '(':
		item = newToken(token.LPAREN, l.char)
	case ')':
		item = newToken(token.RPAREN, l.char)
	case '{':
		item = newToken(token.LBRACE, l.char)
	case '}':
		item = newToken(token.RBRACE, l.char)
	case '\n':
		item = newToken(token.NEWLINE, l.char)
	// no more item to lex
	case 0:
		item.Literal = ""
		item.Type = token.EOF
	// numbers, identifiers, keywords
	default:
		if isLetter(l.char) {
			item.Literal = l.lexIdentifier()
			item.Type = token.LookupIdent(item.Literal)
			return item
		}
		if isDigit(l.char) {
			item.Literal = l.lexNumber()
			item.Type = token.INT
			return item
		}
		item = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return item
}

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '!'
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func newToken(tt token.TokenType, char rune) token.Token {
	return token.Token{Type: tt, Literal: string(char)}
}

func newCompToken(tt token.TokenType, s string) token.Token {
	return token.Token{Type: tt, Literal: s}
}
