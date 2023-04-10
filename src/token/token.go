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
	COLON     = ":"

	EPAREN = "("
	DPAREN = ")"

	EBRACE = "{"
	DBRACE = "}"

	EBRACKET = "["
	DBRACKET = "]"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func VerIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
