package lexer

import "github.com/VivekAlhat/Nova/token"

type Lexer struct {
	input        string
	currPosition int
	nextPosition int
	ch           byte
}

func New(input string) (lexer *Lexer) {
	lexer = &Lexer{input: input}
	lexer.readChar()
	return
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGNMENT, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// identify token type?
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokType token.TokenType, tokValue byte) token.Token {
	return token.Token{Type: tokType, Literal: string(tokValue)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
