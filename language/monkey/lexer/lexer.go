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

	l.skipWhitespace()

	// 現在検査中の文字 l.ch に応じて、token.Token を生成する。
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// 2文字のトークンを生成する。
			ch := l.ch
			l.readChar()
			tk.Literal = string(ch) + string(l.ch)
			tk.Type = token.EQ
		} else {
			tk = newToken(token.ASSIGN, l.ch)
		}
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
	case '-':
		tk = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			// 2文字のトークンを生成する。
			ch := l.ch
			l.readChar()
			tk.Literal = string(ch) + string(l.ch)
			tk.Type = token.NOT_EQ
		} else {
			tk = newToken(token.BANG, l.ch)
		}
	case '/':
		tk = newToken(token.SLASH, l.ch)
	case '*':
		tk = newToken(token.ASTER, l.ch)
	case '<':
		tk = newToken(token.LT, l.ch)
	case '>':
		tk = newToken(token.GT, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case 0:
		// 0 means EOF?
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdent(tk.Literal)

			// readIdentifier の中で既に readChar 進めているので、早期リターンが必要。
			return tk
		} else if isDigit(l.ch) {
			tk.Type = token.INT
			tk.Literal = l.readNumber()

			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	// トークンを返す前に 1 文字読み進める。
	l.readChar()

	return tk
}

// eatWhitespace or consumeWhitespace
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// 1 文字読み進める。
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 影響範囲が大きい。
func isLetter(ch byte) bool {
	// foo_bar のような変数名も許可する。
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 非英字に達するまで読み進める
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		// 1 文字読み進める。
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		// 1 文字読み進める。
		l.readChar()
	}

	return l.input[position:l.position]
}

// readChar との違いは Zl.readPosition を進めないこと。
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL"
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
