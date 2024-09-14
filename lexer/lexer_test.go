package lexer

import (
	"testing"

	"github.com/rubacodes/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;55fn let55 33 *! if(5<10){}else{}return ciao_ \ != ==`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARENT, "("},
		{token.RPARENT, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.INT, "55"},
		{token.FUNCTION, "fn"},
		{token.LET, "let"},
		{token.INT, "55"},
		{token.INT, "33"},
		{token.ASTERISK, "*"},
		{token.BANG, "!"},
		{token.IF, "if"},
		{token.LPARENT, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPARENT, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.RETURN, "return"},
		{token.IDENT, "ciao_"},
		{token.ILLEGAL, "\\"},
		{token.NOT_EQ, "!="},
		{token.EQ, "=="},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Log(tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong: expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {

			t.Fatalf("tests[%d] - tokenLiteral wrong: expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
