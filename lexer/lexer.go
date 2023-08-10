package lexer

import (
	"arthur/token"
	"unicode"
)

type Lexer struct {
	in   string
	pos  int
	read int
	char rune
}

const EOF = -1

// creates a new lexer with the input string and the first character
func NewLexer(input string) *Lexer {
	return &Lexer{
		in:   input,
		pos:  0,
		read: 0,
		char: rune(input[0]),
	}
}

// return the next item in the input string of the one being lexed
func (l *Lexer) peek() rune {
	if l.read >= len(l.in) {
		return EOF
	}
	return rune(l.in[l.read])
}

// advances one position in the input string
func (l *Lexer) next() {
	if l.read >= len(l.in) {
		l.char = EOF
		return
	}
	l.char = rune(l.in[l.read])
	l.pos = l.read
	l.read += 1
}

// goes back one position in the input string
func (l *Lexer) backup() {
	l.read = l.pos
	l.pos -= 1
	l.char = rune(l.in[l.read])
}

// reads keywords and user-defined identifiers
func (l *Lexer) lexIdentifier() string {
	start := l.pos
	for isLetter(l.peek()) {
		l.next()
	}
	str := l.in[start:l.read]
	return str
}

func (l *Lexer) consumeDigits() {
	for isDigit(l.char) {
		l.next()
	}
}

// generates the correct type of token for numbers
func (l *Lexer) lexNumber() token.Token {
	start := l.pos
	var t token.TokenType = token.INT
	l.consumeDigits()

	// checks if it is a float
	if l.char == '.' && !isDigit(l.peek()) {
		l.backup()
	} else if l.char == '.' && isDigit(l.peek()) {
		t = token.FLOAT
		l.next()
		l.consumeDigits()
	}
	return token.Token{Literal: l.in[start:l.read], Type: t}
}

// generates the correct token for the item being lexed
func (l *Lexer) Tokenize() token.Token {
	l.next()

	for isSpace(l.char) {
		l.next()
	}

	switch l.char {
	// operator cases
	case '=':
		if l.peek() == '=' {
			return l.twoCharToken(token.EQ)
		} else {
			return l.token(token.ASSIGN)
		}
	case '+':
		next := l.peek()
		var t token.Token
		if next == '+' {
			t = l.twoCharToken(token.INCREMENT)
		} else if next == '=' {
			t = l.twoCharToken(token.PLUS_EQ)
		} else {
			t = l.token(token.PLUS)
		}
		return t
	case '-':
		next := l.peek()
		var t token.Token
		if next == '-' {
			t = l.twoCharToken(token.DECREMENT)
		} else if next == '=' {
			t = l.twoCharToken(token.MINUS_EQ)
		} else {
			t = l.token(token.MINUS)
		}
		return t
	case '!':
		if l.peek() == '=' {
			return l.twoCharToken(token.NOT_EQ)
		} else {
			return l.token(token.BANG)
		}
	case '*':
		if l.peek() == '=' {
			return l.twoCharToken(token.ASTERISK_EQ)
		} else {
			return l.token(token.ASTERISK)
		}
	case '/':
		if l.peek() == '=' {
			return l.twoCharToken(token.SLASH_EQ)
		} else {
			return l.token(token.SLASH)
		}
	case '<':
		return l.token(token.LT)
	case '>':
		return l.token(token.RT)
	// delimiter cases
	case ',':
		return l.token(token.COMMA)
	case ';':
		return l.token(token.SEMICOLON)
	case '(':
		return l.token(token.LPAREN)
	case ')':
		return l.token(token.RPAREN)
	case '{':
		return l.token(token.LBRACE)
	case '}':
		return l.token(token.RBRACE)
	case '\n':
		return l.token(token.NEWLINE)
	// no more item to lex
	case EOF:
		return token.Token{
			Literal: "",
			Type:    token.EOF,
		}
	// numbers, identifiers, keywords
	default:
		if isLetter(l.char) {
			ident := l.lexIdentifier()
			return token.Token{
				Literal: ident,
				Type:    token.LookupIdent(ident),
			}
		}
		if isDigit(l.char) {
			return l.lexNumber()
		}
		return l.token(token.ILLEGAL)
	}
}

// doesn't skip newline since it replaces semicolons
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isDigit(r rune) bool {
	return unicode.IsNumber(r)
}

// returns a single character token
func (l *Lexer) token(tt token.TokenType) token.Token {
	return token.Token{
		Type:    tt,
		Literal: string(l.in[l.pos]),
	}
}

// returns a two-character token
func (l *Lexer) twoCharToken(tt token.TokenType) token.Token {
	item := token.Token{
		Type:    tt,
		Literal: l.in[l.pos : l.read+1],
	}
	l.next()
	return item
}
