package src

import (
	"fmt"
	"os"
)

func Interpret(expression expr) int {
	return eval(expression)
} 

func eval(expresison expr) int {
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
	fmt.Println("Runtime error")
	os.Exit(1)
	return 0
}

func evalBinary(expression binaryExpr) int {
	left := eval(expression.leftExpr)
	switch expression.operator {
	case "+":
		right := eval(expression.rightExpr)
		return left + right
	
	case "-":
		right := eval(expression.rightExpr)
		return left - right

	case "*":
		right := eval(expression.rightExpr)
		return left * right

	case "/":
		right := eval(expression.rightExpr)
		if right == 0 {
			fmt.Println("Division by zero error")
			os.Exit(1)
			return 0
		}
		return left / right
	}
	fmt.Println("Runtime error")
	os.Exit(1)
	return 0
}

func evalNegand(expression negandExpr) int {
	right := eval(expression.rightExpr)
	return -right
}

func evalGroup(expression groupExpr) int {
	return eval(expression.expression)
}

func evalLiteral(expression literalExpr) int {
	return int(expression)
}

