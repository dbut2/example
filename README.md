# example

example is a Go package designed to bridge the gap between complex business logic or technical implementations and non-technical users. It provides a simple, user-friendly interface for interacting with and testing various functions or services using markdown-formatted tables as input and output.

For examples of this package in use, visit the [examples documentation](https://pkg.go.dev/github.com/dbut2/example/examples).

## Purpose

The main goal of this package is to allow users without programming experience to:
- Easily input data and parameters for testing complex functions or services.
- Visualize the results in a clear, tabular format.
- Experiment with different inputs to understand the underlying logic or behavior of a system.

By using markdown tables, which are human-readable and easy to edit, this package enables a wide range of users to interact with and gain insights from technical implementations without needing to understand or write code.

##  Installation

To use this package in your Go project, you can install it using:

```bash
$ go get github.com/dbut2/example
```

## Usage

The main function provided by this package is `Run`. It takes two parameters:
1. A function that processes the input data and returns the output data
2. A string containing the input data in markdown table format

For practical examples of how to use this package, please check the `/examples` directory in the repository. This directory contains various use cases and implementations that demonstrate how to expose different types of logic or services using markdown tables.

## Contributing

For information on how to contribute to this project, please see the [CONTRIBUTING.md](CONTRIBUTING.md) file.
