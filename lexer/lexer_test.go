package lexer

import (
	"testing"

	"github.com/bamchoh/json_parser/token"
)

func TestNextToken(t *testing.T) {
	input := `[]{},;:`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWithRootIsArray(t *testing.T) {
	input := `[
	1,
	22,
	33,
	"string",
	"abc",
	"",
	"こんにちわ",
	"true",
	true,
	false,
	null
	]`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACKET, "["},
		{token.NUMBER, "1"},
		{token.COMMA, ","},
		{token.NUMBER, "22"},
		{token.COMMA, ","},
		{token.NUMBER, "33"},
		{token.COMMA, ","},
		{token.STRING, "string"},
		{token.COMMA, ","},
		{token.STRING, "abc"},
		{token.COMMA, ","},
		{token.STRING, ""},
		{token.COMMA, ","},
		{token.STRING, "こんにちわ"},
		{token.COMMA, ","},
		{token.STRING, "true"},
		{token.COMMA, ","},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.NULL, "null"},
		{token.RBRACKET, "]"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
