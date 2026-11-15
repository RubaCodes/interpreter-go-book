package parser

import (
	"fmt"

	"github.com/rubacodes/monkey/ast"
	"github.com/rubacodes/monkey/lexer"
	"github.com/rubacodes/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	errors []string

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// read the two tokens, so both curr and peek are set in working wondition
	p.NextToken()
	p.NextToken()
	return p
}
func (p *Parser) NextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// keep parsing until EOF is reached
	for !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}
	return program
}
func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// let
	stmnt := &ast.LetStatement{Token: p.currToken}

	// after let we expext and IDENT
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmnt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
	// after the ident we expct assign...
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO :we're skipping expressions until we meet a semicolon
	for !p.currTokenIs(token.SEMICOLON) {
		p.NextToken()
	}
	return stmnt
}

// Helpers

// check if the curr token satisfies the check , does not advance the parser
func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

// check if the next token satisfies the check , does not advance the parser
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got=%s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// check if the next token satisfies the check, if true advances the parser to the next token
func (p *Parser) expectPeek(tok token.TokenType) bool {
	if p.peekTokenIs(tok) {
		p.NextToken()
		return true
	}
	p.peekError(tok)
	return false
}
