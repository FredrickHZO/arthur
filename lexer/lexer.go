package lexer

import (
	"arthur/token"
)

type Lexer struct {
	in   string
	pos  int
	char rune
}

const EOF = -1

func NewLexer(input string) *Lexer {
	return &Lexer{
		in:   input,
		pos:  0,
		char: rune(input[0]),
	}
}

// return the next item in the input string of the one being lexed
func (l *Lexer) peek() rune {
	if (l.pos + 1) >= len(l.in) {
		return EOF
	}
	return rune(l.in[l.pos+1])
}

// advances one position in the input string
func (l *Lexer) next() {
	l.pos += 1
	if l.pos >= len(l.in) {
		l.char = EOF
		return
	}
	l.char = rune(l.in[l.pos])
}

// reads keywords and user-defined identifiers
func (l *Lexer) lexIdentifier() string {
	startPos := l.pos
	for isLetter(l.char) || isDigit(l.char) {
		l.next()
	}
	// prevents bugs in the case an identifier is the last item
	// to tokenize in the string
	if l.pos >= (len(l.in) - 1) {
		l.pos += 1
	}
	return l.in[startPos:l.pos]
}

// reads a number until there are no digit left
func (l *Lexer) lexNumber() string {
	startPos := l.pos
	for isDigit(l.char) {
		l.next()
	}
	// prevents bugs in the case a number is the last item
	// to tokenize in the string
	if l.pos >= (len(l.in) - 1) {
		l.pos += 1
	}
	return l.in[startPos:l.pos]
}

// doesn't skip newline since it replaces semicolons
func (l *Lexer) skipSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\r' {
		l.next()
	}
}

// generates the correct token for the item being lexed
func (l *Lexer) Tokenize() token.Token {
	var item token.Token

	l.skipSpace()

	switch l.char {
	// operator cases
	case '=':
		if l.peek() == '=' {
			item = l.twoCharToken(token.EQ)
		} else {
			item = l.token(token.ASSIGN)
		}
	case '+':
		next := l.peek()
		if next == '+' {
			item = l.twoCharToken(token.INCREMENT)
		} else if next == '=' {
			item = l.twoCharToken(token.PLUS_EQ)
		} else {
			item = l.token(token.PLUS)
		}
	case '-':
		next := l.peek()
		if next == '-' {
			item = l.twoCharToken(token.DECREMENT)
		} else if next == '=' {
			item = l.twoCharToken(token.MINUS_EQ)
		} else {
			item = l.token(token.MINUS)
		}
	case '!':
		if l.peek() == '=' {
			item = l.twoCharToken(token.NOT_EQ)
		} else {
			item = l.token(token.BANG)
		}
	case '*':
		if l.peek() == '=' {
			item = l.twoCharToken(token.ASTERISK_EQ)
		} else {
			item = l.token(token.ASTERISK)
		}
	case '/':
		if l.peek() == '=' {
			item = l.twoCharToken(token.SLASH_EQ)
		} else {
			item = l.token(token.SLASH)
		}
	case '<':
		item = l.token(token.LT)
	case '>':
		item = l.token(token.RT)
	// delimiter cases
	case ',':
		item = l.token(token.COMMA)
	case ';':
		item = l.token(token.SEMICOLON)
	case '(':
		item = l.token(token.LPAREN)
	case ')':
		item = l.token(token.RPAREN)
	case '{':
		item = l.token(token.LBRACE)
	case '}':
		item = l.token(token.RBRACE)
	case '\n':
		item = l.token(token.NEWLINE)
	// no more item to lex
	case EOF:
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
		item = l.token(token.ILLEGAL)
	}

	l.next()
	return item
}

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '!' || char == '_'
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

// returns a single character token
func (l *Lexer) token(tt token.TokenType) token.Token {
	return token.Token{
		Type:    tt,
		Literal: string(l.char),
	}
}

// returns a two-character token
func (l *Lexer) twoCharToken(tt token.TokenType) token.Token {
	item := token.Token{
		Type:    tt,
		Literal: l.in[l.pos : l.pos+2],
	}
	l.next()
	return item
}
