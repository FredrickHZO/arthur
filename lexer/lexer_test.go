package lexer

import (
	"arthur/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5
		let ten = 10

		let add = fn(x, y) {
			x + y
		}

		let result = add(five, ten)
		
		!-/*5
		5 < 10 > 5
		
		if 5 < 10 {
			return true
		} else {
			return false
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

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
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

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RT, ">"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},

		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
	}

	l := NewLexer(input)

	for i, item := range tests {
		tkn := l.NextToken()

		fmt.Print(tkn.Literal)

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
