package token

// might change to int or byte in the future
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENT      = "IDENT"
	INT        = "INT"
	ASSIGN     = "="
	PLUS       = "+"
	COMMA      = ","
	SEMICOLON  = ";"
	LPAREN     = "("
	RPAREN     = ")"
	LBRACE     = "{"
	RBRACE     = "}"
	WHITESPACE = "\n"
	FUNCTION   = "FUNCTION"
	LET        = "LET"
)
