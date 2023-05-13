package lexer

import "monkey-language/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL"
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// usage of position and readPosition
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	// 現在検査中の文字 l.ch に応じて、token.Token を生成する。
	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		// 0 means EOF?
		tk.Literal = ""
		tk.Type = token.EOF
	}

	// トークンを返す前に 1 文字読み進める。
	l.readChar()

	return tk
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
