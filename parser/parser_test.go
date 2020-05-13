package parser

import (
	"testing"

	"github.com/negli0/shgo/lexer"
)

func TestCmdStatement(t *testing.T) {
	input := `ls -l hoge.go`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"ls"},
		{"-l"},
		{"hoge.go"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if tt.expectedIdentifier != stmt.TokenLiteral() {
			t.Fatalf("tt.expectedIdentifier does not %s. got=%s", tt.expectedIdentifier, stmt.TokenLiteral())
		}
	}
}
