package solver

import (
	"context"
	"strings"
	"testing"
)

func TestProcessorProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	r := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)
	data := []float64{22, 40, 0}
	hasError := []bool{true, true, false}

	for i, d := range data {
		result, err := p.ProcessExpression(context.Background(), r)
		if err != nil && hasError[i] {
			t.Error(err)
		}

		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}
	}
}
