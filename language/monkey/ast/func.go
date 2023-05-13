package ast

import (
	"bytes"
	"monkey-language/token"
	"strings"
)

type FunctionalLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionalLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

func (fl *FunctionalLiteral) expressionNode() {}
func (fl *FunctionalLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
