package token

// might change to int or byte in the future
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TODO: add floats and other symbols
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	NEWLINE   = "NEWLINE"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

// this table is useful to discern between user-defined
// identifiers and keywords
var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
	"\n":  NEWLINE,
}

// checks the table to see if the identifier is
// a keyword of the language or a user-defined one.
// If it's a keyword, it returns its TokenType constant
// else we use IDENT which is the TokenType for
// all user-defined identifiers
func LookupIdent(ident string) TokenType {
	if item, ok := keywords[ident]; ok {
		return item
	}
	return IDENT
}
