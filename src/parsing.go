package src

import (
	"errors"
	"fmt"
)

func Parse(tokens []token) (expr, error) {
	p := parser{tokens, 0}
	ast, err := p.expression()
	return ast, err
}

type expr interface {}

type powerExpr struct {
	base expr
	exponent expr
}

type negandExpr struct {
	expression expr
}

func newNegand(expression expr) negandExpr {
	return negandExpr{expression}
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

func newPowerExpr(base expr, exponent expr) powerExpr {
	return powerExpr{base, exponent}
}

type literalExpr int

// expression -> term 
// term -> factor (("+" | "-") factor)*
// factor -> negand (("/" | "*") negand)*
// negand -> "-" negand | primary
// primary -> "(" expression ")" | number

type parser struct {
	tokens []token
	currentIndex int 
}

func (p *parser) atEnd() bool {
	return p.currentIndex == len(p.tokens)
}

func (p *parser) current() token {
	return p.tokens[p.currentIndex]
}

func (p *parser) advance() {
	p.currentIndex++
}

func (p *parser) match(tokenIds ... int) bool {
	if p.atEnd() {
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
	return p.equality()
	// return p.term()
}

func (p *parser) equality() (expr, error) {
	left, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(EQUAL_EQUAL) {
		p.advance()
		right, err := p.equality()
		if err != nil {
			return nil, err
		}
		left = newbinaryExpr(left, "==", right)
	}
	return left, nil
}

func (p *parser) term() (expr, error) {
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

func (p *parser) factor() (expr, error) {
	// factor -> negand (("/" | "*") negand)*
	expression, err := p.negand()
	if err != nil {
		return nil, err
	}
	for p.match(SLASH, STAR) {
		var operator = p.current().Lexeme
		p.advance()
		// rightExpr := p.negand()
		rightExpr, err := p.negand()
		if err != nil {
			return nil, err
		}
		expression = newbinaryExpr(expression, operator, rightExpr)
	}
	return expression, nil
}

func (p *parser) negand() (expr, error) {
	// negand -> "-" negand | primary
	if p.match(MINUS) {
		// always true if only inputed `-`
		// because advance
		p.advance()
		rightExpr, err := p.negand()
		if err != nil {
			return nil, err
		}
		return newNegand(rightExpr), nil
	}
	return p.power()
	// return p.primary()
}

func (p *parser) power() (expr, error) {
	base, err := p.primary()
	if err != nil {
		return nil, err
	}
	for p.match(POWER) {
		p.advance()
		exponent, err := p.power()
		if err != nil {
			return nil, err
		}
		base = newPowerExpr(base, exponent)
	}
	return base, nil
}

func (p *parser) primary() (expr, error) {
	if p.match(LEFT_PAREN) {
		p.advance()
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		if !p.match(RIGHT_PAREN) {
			return nil, errors.New("Parse Error: Expected token `)`")
		}
		p.advance()
		return newgroupExpr(expression), nil
	} else if p.match(NUM_LITERAL) {
		val := p.current().Value.(int)
		p.advance()
		return literalExpr(val), nil
	}
	if p.atEnd() {
		return nil, errors.New("Parse Error: Expected number literal")
	}
	return nil, fmt.Errorf("Parse Error: Unexpected token `%s`", p.current().Lexeme)
}