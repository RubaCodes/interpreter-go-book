package parser

import (
	"testing"

	"github.com/rubacodes/monkey/ast"
	"github.com/rubacodes/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input :=
		`let x = 5;
		 let y = 10;
		 let foobar = 838383;
		`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)
	// inital structural checks, we have 3 correct statements, we expect 3 statement out
	if len(program.Statements) != 3 {
		t.Fatalf("Program statements doesn't contain 3 statement, got %d", len(program.Statements))
	}

	//table Tests
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// cannot be a let statement if it doesn't start with let...
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() not 'let', got=%q", s.TokenLiteral())
		return false
	}
	// check for proper type
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}
	// check if the value of the identity matches the name given in the code
	if letStmt.Name.Value != name {
		t.Errorf("wrong letStmt.Name.Value: got=%s , want=%s", letStmt.Name.Value, name)
		return false
	}
	if letStmt.Name.TokenLiteral() != name /*ident */ {
		t.Errorf("wrong letStmt.Name.TokenLiteral() : got=%s , want=%s", letStmt.Name.TokenLiteral(), name /*ident */)
		return false
	}
	return true
}
func TestReturnStatements(t *testing.T) {
	input :=
		`return 5;
		 return 10;
		 return 993322;
		`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserError(t, p)
	// inital structural checks, we have 3 correct statements, we expect 3 statement out
	if len(program.Statements) != 3 {
		t.Fatalf("Program statements doesn't contain 3 statement, got %d", len(program.Statements))
	}

	//table Tests

	for _, stmnt := range program.Statements {
		returnStmt, ok := stmnt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmnt not a *ast.ReturnStatement, got=%T", stmnt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("wrong returnStmt.TokenLiteral: want 'return', got=%s", returnStmt.TokenLiteral())
		}
	}
}
func checkParserError(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, e := range errors {
		t.Errorf("parse error: %q", e)
	}
	t.FailNow()
}
