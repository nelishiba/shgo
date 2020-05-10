package lexer

import (
	"testing"

	"github.com/negli0/shgo/token"
)

func TestNextToken(t *testing.T) {
	input := `cat hoge.go | less; for i in a b c d e; do echo hoge; done`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "cat"},
		{token.IDENT, "hoge.go"},
		{token.PIPE, "|"},
		{token.IDENT, "less"},
		{token.SEMICOLON, ";"},
		{token.FOR, "for"},
		{token.IDENT, "i"},
		{token.IN, "in"},
		{token.IDENT, "a"},
		{token.IDENT, "b"},
		{token.IDENT, "c"},
		{token.IDENT, "d"},
		{token.IDENT, "e"},
		{token.SEMICOLON, ";"},
		{token.DO, "do"},
		{token.IDENT, "echo"},
		{token.IDENT, "hoge"},
		{token.SEMICOLON, ";"},
		{token.DONE, "done"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expevted=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
