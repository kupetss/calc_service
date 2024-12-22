package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	var numStack []float64
	var opStack []rune

	applyOperator := func(operator rune, b, a float64) (float64, error) {
		switch operator {
		case '+':
			return a + b, nil
		case '-':
			return a - b, nil
		case '*':
			return a * b, nil
		case '/':
			if b == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return a / b, nil
		default:
			return 0, fmt.Errorf("unknown operator")
		}
	}

	precedence := func(op rune) int {
		switch op {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		}
		return 0
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
				j++
			}
			numStr := expression[i:j]
			number, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number")
			}
			numStack = append(numStack, number)
			i = j - 1
		} else if ch == '(' {
			opStack = append(opStack, ch)
		} else if ch == ')' {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				if len(numStack) < 2 {
					return 0, fmt.Errorf("invalid expression")
				}
				b := numStack[len(numStack)-1]
				a := numStack[len(numStack)-2]
				numStack = numStack[:len(numStack)-2]
				result, err := applyOperator(op, b, a)
				if err != nil {
					return 0, err
				}
				numStack = append(numStack, result)
			}
			if len(opStack) == 0 || opStack[len(opStack)-1] != '(' {
				return 0, fmt.Errorf("invalid expression")
			}
			opStack = opStack[:len(opStack)-1]
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			for len(opStack) > 0 && precedence(opStack[len(opStack)-1]) >= precedence(ch) {
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]
				if len(numStack) < 2 {
					return 0, fmt.Errorf("invalid expression")
				}
				b := numStack[len(numStack)-1]
				a := numStack[len(numStack)-2]
				numStack = numStack[:len(numStack)-2]
				result, err := applyOperator(op, b, a)
				if err != nil {
					return 0, err
				}
				numStack = append(numStack, result)
			}
			opStack = append(opStack, ch)
		} else {
			return 0, fmt.Errorf("invalid character in expression")
		}
	}

	for len(opStack) > 0 {
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		if len(numStack) < 2 {
			return 0, fmt.Errorf("invalid expression")
		}
		b := numStack[len(numStack)-1]
		a := numStack[len(numStack)-2]
		numStack = numStack[:len(numStack)-2]
		result, err := applyOperator(op, b, a)
		if err != nil {
			return 0, err
		}
		numStack = append(numStack, result)
	}

	if len(numStack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return numStack[0], nil
}
