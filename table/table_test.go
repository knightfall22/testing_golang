package table

import (
	"errors"
	"testing"
)

func Test_DoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   error
	}{
		{"addition", 2, 2, "+", 4, nil},
		{"subtraction", 2, 2, "-", 0, nil},
		{"multiplication", 2, 2, "*", 4, nil},
		{"division", 2, 2, "/", 1, nil},
		{"bad_division", 2, 0, "/", 0, DIVISIONERROR},
		{"unknown_op", 2, 0, "%", 0, UNKNOWNERROR},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("expected %d, got %d", d.expected, result)
			}

			if !errors.Is(err, d.errMsg) {
				t.Errorf("expected %v, got %v", d.errMsg, err)
			}
		})
	}

}
