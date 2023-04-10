package lexico

import (
	"fmt"

	"Elysialang/token"
)

type Lexico struct {
	entrada    string //Texto a ser recebido
	posicao    int    // posição atual
	lerPosicao int    // posição atual + 1
	ch         byte   // caractere atual
}

func Novo(entrada string) *Lexico {
	l := &Lexico{entrada: entrada}
	l.lerChar()
	return l
}

func (l *Lexico) ProxToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = novoToken(token.ASSIGN, l.ch)
	case ';':
		tok = novoToken(token.SEMICOLON, l.ch)
	case '(':
		tok = novoToken(token.EPAREN, l.ch)
	case ')':
		tok = novoToken(token.DPAREN, l.ch)
	case ',':
		tok = novoToken(token.COMMA, l.ch)
	case '+':
		tok = novoToken(token.PLUS, l.ch)
	case '{':
		tok = novoToken(token.EBRACE, l.ch)
	case '}':
		tok = novoToken(token.DBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if ehLetra(l.ch) {
			tok.Literal = l.lerIdentificador()
			return tok
		} else if ehDigito(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.lerNumero()
			return tok
		} else {
			tok = novoToken(token.ILLEGAL, l.ch)
		}
	}
	l.lerChar()
	return tok
}

func (l *Lexico) lerChar() {
	if l.lerPosicao >= len(l.entrada) {
		l.ch = 0
	} else {
		l.ch = l.entrada[l.lerPosicao]
	}
	l.posicao = l.lerPosicao
	l.lerPosicao += 1
}
func novoToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexico) lerIdentificador() string {
	posicao := l.posicao
	for ehLetra(l.ch) {
		l.lerChar()
	}
	return l.entrada[posicao:l.posicao]
}

func ehDigito(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func ehLetra(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

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

func (l *Lexico) lerNumero() string {
	prim_posicao := l.posicao
	var numero string
	var pos_meio int
	var eu_tenho_dot bool

	for ehDigito(l.ch) {
		if l.ch == '.' {
			numero = l.entrada[prim_posicao:l.posicao]
			numero = fmt.Sprintf("%s", numero)
			pos_meio = l.posicao
			eu_tenho_dot = true
		}
		l.lerChar()
	}

	floatin := fmt.Sprintf("%s%s", numero, l.entrada[pos_meio:l.posicao])

	if eu_tenho_dot {
		return floatin
	} else {
		return l.entrada[prim_posicao:l.posicao]
	}

}
