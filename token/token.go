package token

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	NUMBER    = "NUMBER"
	STRING    = "STRING"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	NULL      = "NULL"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
)

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(lit string) TokenType {
	if tok, ok := keywords[lit]; ok {
		return tok
	}
	return IDENT
}
