package table

import (
	"errors"
	"fmt"
)

var DIVISIONERROR = errors.New("division by zero")
var UNKNOWNERROR = errors.New("unknown operator")

func DoMath(num1, num2 int, op string) (int, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, DIVISIONERROR
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("%w %s", UNKNOWNERROR, op)
	}
}
