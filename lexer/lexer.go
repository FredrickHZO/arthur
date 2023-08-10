package lexer

import (
	"arthur/token"
	"strings"
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	items chan token.Token
	input string
	start int
	pos   int
	width int
}

type stateFn func(*lexer) stateFn

const EOF = -1

// initializes the lexer and starts to lex the input string
func Lex(in string) chan token.Token {
	l := &lexer{
		input: in,
		items: make(chan token.Token),
	}
	go l.run()
	return l.items
}

// effectively starts lexing the input string
func (l *lexer) run() {
	for state := lexExpr; state != nil; {
		state = state(l)
	}
}

// emit passes an item back to the client
func (l *lexer) emit(t token.TokenType) {
	l.items <- token.Token{
		Type:    t,
		Literal: l.input[l.start:l.pos],
	}
	l.start = l.pos
}

// returns the next char
// doesn't advance in the input string
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// advances one position in the input string
func (l *lexer) next() rune {
	var r rune
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// goes back one position in the input string
func (l *lexer) backup() {
	l.pos -= l.width
}

// skips over the pending input before this point
func (l *lexer) ignore() {
	l.start = l.pos
}

// consumes the next rune if it's from a valid set
func (l *lexer) accept(set string) bool {
	if strings.IndexRune(set, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// consumes a run of runes from the valid set
func (l *lexer) acceptRun(set string) {
	for strings.IndexRune(set, l.next()) >= 0 {
		// goes forward until it finds an invalid item
	}
	l.backup()
}

// newline character must not be recognized as space
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

// returns the current range from start to pos as a string
func (l *lexer) current() string {
	return l.input[l.start:l.pos]
}

// generates a token for a keyword or user-defined identifier
func lexIdentifier(l *lexer) stateFn {
	digits := "abcdefghilmnopqrstuvzkjyxABCDEFGHILMNOPQRSTUVZKJYX123456789_"
	l.acceptRun(digits)
	l.emit(token.LookupIdent(l.current()))
	return lexExpr
}

// returns true if the current character is a letter
func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// generates the correct token for the item being lexed
func lexExpr(l *lexer) stateFn {
	r := l.next()

	switch {
	case r == EOF:
		l.emit(token.EOF)
		return nil

	case r == '\n':
		l.emit(token.NEWLINE)

	case isSpace(r):
		l.ignore()

	case isLetter(r):
		return lexIdentifier
	}

	return lexExpr
}
