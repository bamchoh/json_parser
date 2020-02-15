package parser

import (
	"fmt"
	"strconv"

	"github.com/bamchoh/json_parser/ast"
	"github.com/bamchoh/json_parser/lexer"
	"github.com/bamchoh/json_parser/token"
)

var (
	NULL  = &ast.Null{}
	TRUE  = &ast.Boolean{Value: true}
	FALSE = &ast.Boolean{Value: false}
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Parse() interface{} {
	return p.parse()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) AddError(format string, args ...interface{}) {
	p.errors = append(p.errors, fmt.Sprintf(format, args...))
}

func (p *Parser) parse() interface{} {
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseHash()
	case token.LBRACKET:
		return p.parseArray()
	case token.NUMBER:
		num, _ := strconv.ParseFloat(p.curToken.Literal, 64)
		return num
	case token.STRING:
		return p.curToken.Literal
	case token.TRUE:
		return true
	case token.FALSE:
		return false
	case token.NULL:
		return nil
	}
	return nil
}

func (p *Parser) parseHash() map[string]interface{} {
	data := make(map[string]interface{}, 0)
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key, ok := p.parse().(string)
		if !ok {
			p.AddError("key is not *ast.String got=%T", key)
			return nil
		}

		if !p.peekTokenIs(token.COLON) {
			p.AddError("separator is not COLON, got=%T", p.peekToken)
			return nil
		}
		p.nextToken()
		p.nextToken()
		val := p.parse()

		data[key] = val

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return data
}

func (p *Parser) parseArray() []interface{} {
	data := make([]interface{}, 0)
	for !p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		val := p.parse()

		data = append(data, val)

		if !p.peekTokenIs(token.RBRACKET) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return data
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken, _ = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}
