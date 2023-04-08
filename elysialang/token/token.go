package token

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identificadores
	IDENT = "IDENT"
	INT   = "INT"

	//Delimitadores
	COMMA     = ","
	SEMICOLON = ";"

	EPAREN = "("
	DPAREN = ")"

	EBRACE = "{"
	DBRACE = "}"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
