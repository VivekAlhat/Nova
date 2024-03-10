package lexer

import (
	"testing"

	"github.com/VivekAlhat/Nova/token"
)

func TestNextToken(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		input := `=+(){};
			!-/*5;
			5 < 10 > 5;
		`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.ASSIGNMENT, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.GT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - incorrect token type. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - incorrect literal. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})

	t.Run("More complex test", func(t *testing.T) {
		input := `def one = 1;
			def two = 2;

			def add = fn(a, b) {
				a + b;
			}

			def result = add(one, two);
		`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.DEF, "def"},
			{token.IDENTIFIER, "one"},
			{token.ASSIGNMENT, "="},
			{token.INT, "1"},
			{token.SEMICOLON, ";"},
			{token.DEF, "def"},
			{token.IDENTIFIER, "two"},
			{token.ASSIGNMENT, "="},
			{token.INT, "2"},
			{token.SEMICOLON, ";"},
			{token.DEF, "def"},
			{token.IDENTIFIER, "add"},
			{token.ASSIGNMENT, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENTIFIER, "a"},
			{token.COMMA, ","},
			{token.IDENTIFIER, "b"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENTIFIER, "a"},
			{token.PLUS, "+"},
			{token.IDENTIFIER, "b"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.DEF, "def"},
			{token.IDENTIFIER, "result"},
			{token.ASSIGNMENT, "="},
			{token.IDENTIFIER, "add"},
			{token.LPAREN, "("},
			{token.IDENTIFIER, "one"},
			{token.COMMA, ","},
			{token.IDENTIFIER, "two"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - incorrect token type. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - incorrect literal. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})
}
