package lexer

import "github.com/ericsolomon/pulsar/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var t token.Token

	lexer.eatWhitespace()
	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			t = token.Token{Type: token.EQUAL, Literal: "=="}
		} else {
			t = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		t = newToken(token.PLUS, lexer.ch)
	case ',':
		t = newToken(token.COMMA, lexer.ch)
	case ';':
		t = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		t = newToken(token.LPAREN, lexer.ch)
	case ')':
		t = newToken(token.RPAREN, lexer.ch)
	case '{':
		t = newToken(token.LBRACE, lexer.ch)
	case '}':
		t = newToken(token.RBRACE, lexer.ch)
	case '<':
		t = newToken(token.LT, lexer.ch)
	case '>':
		t = newToken(token.GT, lexer.ch)
	case '*':
		t = newToken(token.ASTERISK, lexer.ch)
	case '/':
		t = newToken(token.FSLASH, lexer.ch)
	case '-':
		t = newToken(token.MINUS, lexer.ch)
	case '!':
		if lexer.peekChar() == '=' {
			lexer.readChar()
			t = token.Token{Type: token.NOTEQUAL, Literal: "!="}
		} else {
			t = newToken(token.BANG, lexer.ch)
		}
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			t.Literal = lexer.readIdent()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(lexer.ch) {
			t.Literal = lexer.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) peekChar() byte {
	if len(lexer.input) >= lexer.readPosition {
		return lexer.input[lexer.readPosition]
	}

	return 0
}

func (lexer *Lexer) eatWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch < '9'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readIdent() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch < 'z' || 'A' <= ch && ch < 'Z' || ch == '_'
}
