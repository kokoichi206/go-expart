package ast

import (
	"monkey-language/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "pien"},
					Value: "pien",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foobar"},
					Value: "foobar",
				},
			},
		},
	}

	if program.String() != "let pien = foobar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
