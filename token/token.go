package token

// might change to int or byte in the future
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TODO: add floats, hex, octal and other operators
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

// keeps track of the language keywords
var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
	"\n":  NEWLINE,
}

// checks if the identifier is present in the keywords table
// returns the correct TokenType if present, otherwise IDENT
func LookupIdent(ident string) TokenType {
	if item, ok := keywords[ident]; ok {
		return item
	}
	return IDENT
}
