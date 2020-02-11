package token

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
