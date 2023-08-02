package lexer

import (
	"arthur/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;\n"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.NEWLINE, "\n"},
	}

	l := NewLexer(input)

	for i, item := range tests {
		tkn := l.NextToken()

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
