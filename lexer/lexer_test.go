package lexer

import (
	"testing"

	"github.com/bamchoh/json_parser/token"
)

func TestNextToken(t *testing.T) {
	input := `{
		"string": 1234,
		"minus": -1234,
		"float": -1234.1234,
		"scientific1": -1234e10,
		"scientific2": -1234e+10,
		"scientific3": -1234e-10,
		"number": -1234.1234e10,
		"": 2345,
		"array": [
			1,2,3,
			true,
			false,
			null,
			"value"
		],
		"escape": "\"\\\/\b\f\n\r\t\u3042\uD83D\ude28"
	}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "string"},
		{token.COLON, ":"},
		{token.NUMBER, "1234"},
		{token.COMMA, ","},
		{token.STRING, "minus"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234"},
		{token.COMMA, ","},
		{token.STRING, "float"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234.1234"},
		{token.COMMA, ","},
		{token.STRING, "scientific1"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234e10"},
		{token.COMMA, ","},
		{token.STRING, "scientific2"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234e+10"},
		{token.COMMA, ","},
		{token.STRING, "scientific3"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234e-10"},
		{token.COMMA, ","},
		{token.STRING, "number"},
		{token.COLON, ":"},
		{token.NUMBER, "-1234.1234e10"},
		{token.COMMA, ","},
		{token.STRING, ""},
		{token.COLON, ":"},
		{token.NUMBER, "2345"},
		{token.COMMA, ","},
		{token.STRING, "array"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.NUMBER, "1"},
		{token.COMMA, ","},
		{token.NUMBER, "2"},
		{token.COMMA, ","},
		{token.NUMBER, "3"},
		{token.COMMA, ","},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.NULL, "null"},
		{token.COMMA, ","},
		{token.STRING, "value"},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.STRING, "escape"},
		{token.COLON, ":"},
		{token.STRING, "\"\\/\b\f\n\r\t\u3042\U0001F628"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok, _ := l.NextToken()

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
	22,33,
	"string",
	"abc","",
	"こんにちわ",
	"true",
	true,false,null,
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
		tok, _ := l.NextToken()

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
