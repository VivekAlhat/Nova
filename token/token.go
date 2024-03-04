package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers
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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"def": DEF,
}

func CheckIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}
