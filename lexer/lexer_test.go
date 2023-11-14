package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `A AND NOT (B OR C)`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "A"},
		{token.AND, "AND"},
		{token.NOT, "NOT"},
		{token.LPAREN, "("},
		{token.IDENT, "B"},
		{token.OR, "OR"},
		{token.IDENT, "C"},
		{token.RPAREN, ")"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
