package lexer

import (
	"arthur/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let num = 5.54

		let add = fn(x, y) {
			x + y
		}

		let result = add(five, 5)
		pippo`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "num"},
		{token.ASSIGN, "="},
		{token.FLOAT, "5.54"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "pippo"},
	}

	l := NewLexer(input)

	for i, item := range tests {
		tkn := l.Tokenize()

		fmt.Println(tkn)

		if tkn.Type != item.expectedType {
			t.Fatalf("tests[%d] - wrong token type, expected=%q, got=%q",
				i, item.expectedType, tkn.Type)
		}

		if tkn.Literal != item.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal type, expected=%q, got =%q",
				i, item.expectedLiteral, tkn.Literal)
		}
	}
}
