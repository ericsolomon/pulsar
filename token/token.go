package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	EQUAL     = "=="
	NOTEQUAL  = "!="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LT       = "<"
	GT       = ">"
	ASTERISK = "*"
	FSLASH   = "/"
	MINUS    = "-"
	BANG     = "!"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	ELSE     = "ELSE"
	IF       = "IF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}

	return IDENT
}
