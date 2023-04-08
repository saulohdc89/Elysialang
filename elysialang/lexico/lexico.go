package lexico

import (
	"github.com/saulohdc89/PROJETOCOMPIL/elysialang/token"
)

type Lexico struct {
	entrada    string
	posicao    int  // atual posição na entrada
	lerPosicao int  // posição atual de leitura na entrada (após o char atual)
	ch         byte // atual char
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
		if isLetra(l.ch) {
			tok.Literal = l.lerIdentificador()
			return tok
		} else if isDigito(l.ch) {
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
	for isLetra(l.ch) {
		l.lerChar()
	}
	return l.entrada[posicao:l.posicao]
}

func isLetra(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
