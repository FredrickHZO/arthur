package lexer

import (
	"arthur/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `foo
		let`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "foo"},
		{token.NEWLINE, "\n"},
		{token.LET, "let"},
		{token.EOF, ""},
	}

	items := Lex(input)

	for i, item := range tests {
		token := <-items

		if token.Type != item.expectedType {
			t.Fatalf("tests[%d] - wrong token type, expected=%q, got=%q",
				i, item.expectedType, token.Type)
		}

		if token.Literal != item.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal type, expected=%q, got =%q",
				i, item.expectedLiteral, token.Literal)
		}
	}
}
