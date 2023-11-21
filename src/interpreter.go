package src

import (
	"errors"
)

func Interpret(expression expr) (float64, error) {
	return eval(expression)
} 

func eval(expresison expr) (float64, error) {
	_, isBin := expresison.(binaryExpr)
	if isBin {
		return evalBinary(expresison.(binaryExpr))
	} 
	_, isnegand := expresison.(negandExpr)
	if isnegand {
		return evalNegand(expresison.(negandExpr))
	}
	_, isGroup := expresison.(groupExpr)
	if isGroup {
		return evalGroup(expresison.(groupExpr))
	}
	_, isLiteral := expresison.(literalExpr)	
	if isLiteral {
		return evalLiteral(expresison.(literalExpr))
	}
	return 0, errors.New("Runtime Error: Unreachable")
}

func evalBinary(expression binaryExpr) (float64, error) {
	left, err := eval(expression.leftExpr)
	if err != nil {
		return 0, err
	}
	switch expression.operator {
	case "+":
		right, err := eval(expression.rightExpr)
		return left + right, err
	
	case "-":
		right, err := eval(expression.rightExpr)
		return left - right, err

	case "*":
		right, err := eval(expression.rightExpr)
		return left * right, err

	case "/":
		right, err := eval(expression.rightExpr)
		if right == 0 {
			return 0, errors.New("Runtime Error: Division by zero error")
		}
		return left / right, err
	}
	return 0, errors.New("Runtime Error: Unsupported operator")
}

func evalNegand(expression negandExpr) (float64, error) {
	rightExp, err := eval(expression.expression)
	return -rightExp, err
}

func evalGroup(expression groupExpr) (float64, error) {
	return eval(expression.expression)
}

func evalLiteral(expression literalExpr) (float64, error) {
	return float64(expression), nil
}

