package calculator_test

import (
	"github.com/dbut2/example"
	"github.com/dbut2/example/examples/calculator"
)

// Play With Me
var maxInput = `
| Number |
|--------|
| 5      |
| 6      |
| 2      |
| 4      |
| 1      |
| 3      |
`

func ExampleMaximum() {
	example.Run(calculator.Maximum, maxInput)
	// Output:
	// | Number |
	// |--------|
	// | 6      |
}

// Play With Me
var calculateInput = `
| Operand | Number |
|---------|--------|
| +       | 10     |
| -       | 5      |
| *       | 4      |
| /       | 2      |
`

func ExampleCalculate() {
	example.Run(calculator.Calculate, calculateInput)
	// Output:
	// | Result |
	// |--------|
	// | 10     |
}
