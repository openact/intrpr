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

	IF    = "IF"
	ELSE  = "ELSE"
	THEN  = "THEN"
	EQ    = "="
	PLUS  = "+"
	MINUS = "-"
	LT    = "<"
	GT    = ">"
	DIV   = "/"
	MUL   = "*"

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
	"AND":   AND,
	"and":   AND,
	"And":   AND,
	"OR":    OR,
	"Or":    OR,
	"or":    OR,
	"NOT":   NOT,
	"Not":   NOT,
	"not":   NOT,
	"IF":    IF,
	"ELSE":  ELSE,
	"THEN":  THEN,
	"EQ":    EQ,
	"PLUS":  PLUS,
	"MINUS": MINUS,
	"LT":    LT,
	"GT":    GT,
	"DIV":   DIV,
	"MUL":   MUL,
}

var aliases = map[string]string{
	"AND": "&&",
	"And": "&&",
	"and": "&&",
	"OR":  "||",
	"Or":  "||",
	"or":  "||",
	"NOT": "!",
	"Not": "!",
	"not": "!",
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
