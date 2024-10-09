package calculator_test

import (
	"github.com/dbut2/example"
	"github.com/dbut2/example/examples/calculator"
)

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
	// output:
	// | Number |
	// |--------|
	// | 6      |
}

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
	// output:
	// | Result |
	// |--------|
	// | 10     |
}
