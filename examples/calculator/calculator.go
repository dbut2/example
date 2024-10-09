package calculator

import (
	"fmt"
)

type Number struct {
	Number float64
}

func Maximum(numbers []Number) (Number, error) {
	if len(numbers) == 0 {
		return Number{}, fmt.Errorf("no numbers provided")
	}
	m := numbers[0]
	for _, v := range numbers {
		if v.Number > m.Number {
			m = v
		}
	}
	return m, nil
}

type Calculation struct {
	Operand string
	Number  float64
}

type Result struct {
	Result float64
}

func Calculate(c []Calculation) (Result, error) {
	var result float64
	for _, v := range c {
		switch v.Operand {
		case "+":
			result += v.Number
		case "-":
			result -= v.Number
		case "*":
			result *= v.Number
		case "/":
			result /= v.Number
		}
	}
	return Result{Result: result}, nil
}
