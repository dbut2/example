// Package calculator provides basic arithmetic operations on numeric data.
package calculator

import (
	"fmt"
)

// Number represents a single numeric value.
type Number struct {
	Number float64
}

// Maximum finds the largest number in a slice of Numbers.
// It returns an error if the input slice is empty.
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

// Calculation represents a single step in a series of arithmetic operations.
type Calculation struct {
	Operand string
	Number  float64
}

// Result represents the outcome of a series of calculations.
type Result struct {
	Result float64
}

// Calculate performs a series of arithmetic operations based on the input Calculations.
// It applies the operations in the order they are provided.
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
