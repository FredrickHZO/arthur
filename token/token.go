package token

// might change to int or byte in the future
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TODO: add floats, hex, octal and other operators
const (
	// identifiers & literals
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"
	// operators
	EQ       = "=="
	NOT_EQ   = "!="
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	RT       = ">"
	// delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	NEWLINE   = "NEWLINE"
	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

// keeps track of the language keywords
var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"\n":     NEWLINE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

// checks if the identifier is present in the keywords table
// returns the correct TokenType if present, otherwise IDENT
func LookupIdent(ident string) TokenType {
	if item, ok := keywords[ident]; ok {
		return item
	}
	return IDENT
}
