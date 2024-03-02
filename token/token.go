package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// operators
	ASSIGNMENT = "="
	PLUS       = "+"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keywords
	DEF      = "DEF"
	FUNCTION = "FUNCTION"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
