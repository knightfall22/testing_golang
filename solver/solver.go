package solver

import (
	"context"
	"errors"
	"fmt"
	"io"
)

type Processor struct {
	Solver MathSolver
}

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) {
	curExpression, err := readToNewLine(r)

	if err != nil {
		return 0, err
	}

	if len(curExpression) == 0 {
		return 0, fmt.Errorf("no expression here")
	}

	answer, err := p.Solver.Resolve(ctx, curExpression)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}
