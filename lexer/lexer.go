package lexer

import (
	"github.com/VivekAlhat/Nova/token"
)

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

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currPosition]
}

func (l *Lexer) ignoreWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.ignoreWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = makeTwoCharToken(l.ch, token.EQ, l)
		} else {
			tok = newToken(token.ASSIGNMENT, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = makeTwoCharToken(l.ch, token.NOTEQ, l)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
			tok.Type = token.CheckIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
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

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func makeTwoCharToken(curr byte, tokType token.TokenType, l *Lexer) (tok token.Token) {
	ch := curr
	l.readChar()
	tok = token.Token{Type: tokType, Literal: string(ch) + string(l.ch)}
	return tok
}
