package lexico

import (
	"testing"

	"Elysialang/token"
)

func TesteProximoToken(t *testing.T) {
	entrada := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);
	`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.EPAREN, "(=)"},
		{token.DPAREN, "="},
		{token.EBRACE, "="},
		{token.DBRACE, "="},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
		{token.FUNC, "fn"},
	}

	l := Novo(entrada)

	for i, tt := range tests {
		tok := l.ProxToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("teste[%d] tipo de token errado, esperado = %q, teve=%q",
				i, &tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectLiteral {
			t.Fatalf("teste[%d] - literal errado, esperado =%q, teve =%q", i, tt.expectLiteral, tok.Literal)

		}
	}
}
