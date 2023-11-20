package src

import (
	"errors"
)

func Parse(tokens []token) (expr, error) {
	p := parser{tokens, 0}
	ast, err := p.expression()
	return ast, err
}

type expr interface {}

type unaryExpr struct {
	operator string
	rightExpr expr
}

func newUnary(op string, expression expr) unaryExpr {
	return unaryExpr{op, expression}
}

type binaryExpr struct {
	leftExpr expr
	operator string
	rightExpr expr	
}

func newbinaryExpr(leftExpr expr, op string, rightExpr expr) binaryExpr {
	return binaryExpr{leftExpr, op, rightExpr}
}

type groupExpr struct {
	expression expr
}

func newgroupExpr(expression expr) groupExpr {
	return groupExpr{expression}
}

type literalExpr int

// expression -> term 
// term -> factor (("+" | "-") factor)*
// factor -> unary (("/" | "*") unary)*
// unary -> "-" unary | primary
// primary -> "(" expressiion ")" | number

type parser struct {
	tokens []token
	currentIndex int 
}

func (p *parser) current() token {
	return p.tokens[p.currentIndex]
}

func (p *parser) advance() {
	if p.currentIndex < len(p.tokens) {
		p.currentIndex++
	}
}

func (p *parser) match(tokenIds ... int) bool {
	if p.currentIndex >= len(p.tokens) {
		return false
	}
	for _, t := range tokenIds {
		if p.current().TokenId == t {
			return true
		}
	}
	return false
}

func (p *parser) expression() (expr, error) {
	return p.term()
}

func (p *parser)term() (expr, error) {
	expression, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.match(PLUS, MINUS) {
		var operator = p.current().Lexeme
		p.advance()
		rightExpr, err := p.factor()
		if err != nil {
			return nil, err
		}
		expression = newbinaryExpr(expression, operator, rightExpr)
	}
	return expression, nil
}

func (p *parser)factor() (expr, error) {
	// factor -> unary (("/" | "*") unary)*
	expression, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(SLASH, STAR) {
		var operator = p.current().Lexeme
		p.advance()
		// rightExpr := p.unary()
		rightExpr, err := p.unary()
		if err != nil {
			return nil, err
		}
		expression = newbinaryExpr(expression, operator, rightExpr)
	}
	return expression, nil
}

func (p *parser)unary() (expr, error) {
	// unary -> "-" unary | primary
	for p.match(MINUS) {
		var operator = p.current().Lexeme
		p.advance()
		// rightExpr := p.unary()
		rightExpr, err := p.unary()
		if err != nil {
			return nil, err
		}
		return newUnary(operator, rightExpr), nil
	}
	return p.primary()
}

func (p *parser)primary() (expr, error) {
	if p.match(LEFT_PAREN) {
		p.advance()
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		if !p.match(RIGHT_PAREN) {
			return nil, errors.New("Expected Token ')'")
		}
		return newgroupExpr(expression), nil
	} else if p.match(NUM_LITERAL) {
		val := p.current().Value.(int)
		p.advance()
		return literalExpr(val), nil
	}
	return nil, errors.New("Unexpected Token")
}