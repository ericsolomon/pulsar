package lexer

import (
	"github.com/ericsolomon/pulsar/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	testInput := `
let x = 5;
let y = 10;

let add = fn(x,y) {
  x + y;
};

let result = add(x, y);

!-/*5t;
x < 5 > y;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
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
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.FSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.IDENT, "t"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.INT, "5"},
		{token.GT, ">"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	testLexer := New(testInput)
	for i, test := range tests {
		token := testLexer.nextToken()

		if token.Type != test.expectedType {
			t.Fatalf("test %d failed. Expected token type %q, got %q", i, test.expectedType, token.Type)
		}

		if token.Literal != test.expectedLiteral {
			t.Fatalf("test %d failed. Expected literal %q, got %q", i, test.expectedLiteral, token.Literal)
		}
	}
}
