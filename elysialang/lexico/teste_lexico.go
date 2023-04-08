package lexico

import (
	"testing"

	"github.com/saulohdc89/PROJETOCOMPIL/elysialang/token"
)

func TesteProximoToken(t *testing.T) {
	input := `=-(){},;`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "="},
		{token.EPAREN, "="},
		{token.DPAREN, "="},
		{token.EBRACE, "="},
		{token.DBRACE, "="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := Novo(input)
	for i, tt := range tests {
		tok := l.ProxToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("teste[%d] tipo de token errado, esperado = %q, teve=%q",
				i, &tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectLiteral {
			t.Fatalf("teste[%d] - literal errado, expected =%q, got =%q", i, tt.expectLiteral, tok.Literal)

		}
	}
}
