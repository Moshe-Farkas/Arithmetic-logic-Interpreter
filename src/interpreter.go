package src

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func Interpret(expression expr) (any, error) {
	return eval(expression)
} 

func eval(expresison expr) (any, error) {
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
	_, isBoolean := expresison.(BooleanExpr)
	if isBoolean {
		return evalBoolean(expresison.(BooleanExpr))
	}
	return 0, errors.New("Runtime Error: Unreachable")
}

func evalBinary(expression binaryExpr) (any, error) {
	left, err := eval(expression.leftExpr)
	if err != nil {
		return 0, err
	}
	switch expression.operator {
	case "+":
		right, err := eval(expression.rightExpr)
		if checkNumOperands(left, right) {
			return left.(float64) + right.(float64), err
		} else {
			return nil, fmt.Errorf("Runtime Error: cannot add opperand of type `%T` to type `%T`", left, right)
		}
	
	case "-":
		right, err := eval(expression.rightExpr)
		if checkNumOperands(left, right) {
			return left.(float64) - right.(float64), err
		} else {
			return nil, fmt.Errorf("Runtime Error: cannot subtract opperand of type `%T` to type `%T`", left, right)
		}

	case "*":
		right, err := eval(expression.rightExpr)
		if checkNumOperands(left, right) {
			return left.(float64) * right.(float64), err
		} 
		return nil, fmt.Errorf("Runtime Error: cannot multiply opperand of type `%T` to type `%T`", left, right)

	case "==":
		right, err := eval(expression.rightExpr)
		if err != nil {
			return nil, err
		}
		return equal(left, right)

	case "/":
		right, err := eval(expression.rightExpr)
		if checkNumOperands(left, right) {
			if right.(float64) == 0 {
				return 0, errors.New("Runtime Error: Division by zero")
			}
			return left.(float64) / right.(float64), err
		} else {
			return nil, fmt.Errorf("Runtime Error: cannot divide opperand of type `%T` to type `%T`", left, right)
		}

	case "^":
		right, err := eval(expression.rightExpr)
		if checkNumOperands(left, right) {
			return math.Pow(left.(float64), right.(float64)), err
		}
		return nil, fmt.Errorf("Runtime Error: cannot raise opperand of type `%T` to type `%T`", left, right)
	
	case ">", "<", ">=", "<=":
		right, err := eval(expression.rightExpr)
		if err != nil {
			return false, err
		}
		if !checkNumOperands(left, right) {
			return false, fmt.Errorf("Runtime Error: cannot compare opperand of type `%T` to type `%T`", left, right)
		}
		return comp(left.(float64), right.(float64), expression.operator)

	// case ">":
	// 	right, err := eval(expression.rightExpr)
	// 	if checkNumOperands(left, right) {
	// 		return left.(float64) > right.(float64), err
	// 	} 
	// 	return nil, fmt.Errorf("Runtime Error: cannot compare opperand of type `%T` to type `%T`", left, right)
	// case "<":
	// 	right, err := eval(expression.rightExpr)
	// 	if checkNumOperands(left, right) {
	// 		return left.(float64) < right.(float64), err
	// 	} 
	// 	return nil, fmt.Errorf("Runtime Error: cannot compare opperand of type `%T` to type `%T`", left, right)
	}
	return nil, errors.New("Runtime Error: Unsupported operator")
}

func checkNumOperands(left, right expr) bool {
	return reflect.TypeOf(left) == reflect.TypeOf(0.0) && 
		   reflect.TypeOf(right) == reflect.TypeOf(0.0)
}

func evalNegand(expression negandExpr) (any, error) {
	rightExp, err := eval(expression.expression)
	return -(rightExp.(float64)), err
}

func evalGroup(expression groupExpr) (any, error) {
	return eval(expression.expression)
}

func evalLiteral(expression literalExpr) (any, error) {
	return float64(expression), nil
}

func evalBoolean(expression BooleanExpr) (any, error) {
	return bool(expression), nil
}

func equal(left, right expr) (bool, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return false, fmt.Errorf("Runtime Error: cannot compare opperand of type `%T` to type `%T`", left, right)
	}
	return reflect.DeepEqual(left, right), nil
}

func comp(left, right float64, operator string) (bool, error) {
	switch operator {
	case ">":
		return left > right, nil
	case "<":
		return left < right, nil
	case ">=":
		return left >= right, nil
	case "<=":
		return left <= right, nil
	}
	return false, errors.New("Runtime Error: Unknown comparison operator")
}