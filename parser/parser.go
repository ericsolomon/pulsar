package parser

import (
	"github.com/ericsolomon/pulsar/ast"
	"github.com/ericsolomon/pulsar/lexer"
	"github.com/ericsolomon/pulsar/token"
)

type Parser struct {
	lexer     *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}

	// set curToken and peekToken
	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.curToken.Type != token.EOF {
		statement := parser.parseLetStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: parser.curToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO skipping expressions
	for parser.curToken.Type != token.SEMICOLON {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) expectPeek(tokenType token.TokenType) bool {
	if parser.peekToken.Type == tokenType {
		parser.nextToken()
		return true
	}

	return false
}
