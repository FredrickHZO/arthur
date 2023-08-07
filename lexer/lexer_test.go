package lexer

import (
	"arthur/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5

		let add = fn(x, y) {
			x + y
		}

		let result = add(five, ten)
		
		if !(10 != 9) {
			return false
		}
		if 9 == 9 {
			return true
		}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
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
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		{token.IF, "if"},
		{token.BANG, "!"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.IF, "if"},
		{token.INT, "9"},
		{token.EQ, "=="},
		{token.INT, "9"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, item := range tests {
		tkn := l.Tokenize()

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
