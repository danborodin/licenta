package lexer

import (
	"bdlang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	var two = 2;
	var one = 1;

	var add = func(x, y) {
		x + y;
	}

	var result = add(two, one);
	!-/*7;
	3 < 11 > 3;

	if (2 < 7) {
		return true;
	} else {
		return false;
	}

	9 == 9;
	9 != 6;

	"teststring"
	"test string"
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "one"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "func"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "two"},
		{token.COMMA, ","},
		{token.IDENT, "one"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.LT, "<"},
		{token.INT, "11"},
		{token.GT, ">"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.LT, "<"},
		{token.INT, "7"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "9"},
		{token.EQ, "=="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "9"},
		{token.NOT_EQ, "!="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.STRING, "teststring"},
		{token.STRING, "test string"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected:%q, got:`%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=:%q, got:%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
