package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	LPAREN = "("
	RPAREN = ")"

	AND = "AND"
	OR  = "OR"
	NOT = "NOT"

	SEMICOLON  = ";"
	WHITESPACE = " "
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Alias   string
}

var keywords = map[string]TokenType{
	"AND": AND,
	"OR":  OR,
	"NOT": NOT,
}

var aliases = map[string]string{
	"AND": "&&",
	"OR":  "||",
	"NOT": "!",
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func LookupAlias(alias string) string {
	if tok, ok := aliases[alias]; ok {
		return tok
	}
	return alias
}
