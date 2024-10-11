// Package example provides a generic utility for processing structured data
// using markdown-formatted tables as input and output. It is designed to
// simplify the testing and demonstration of functions that operate on
// tabular data by allowing easy input specification and result visualization.
package example

import (
	"fmt"
	"reflect"
)

// Run executes a function f with input data from a markdown table and outputs the results as a markdown table.
//
// The function f can accept either a single item or a slice/array of items as input (type T),
// and should return a single item or a slice/array of items as output (type U), along with an error.
//
// Run performs the following steps:
// 1. Parses the input markdown table.
// 2. Processes the input data based on the type of T (single item or slice/array).
// 3. Executes the provided function f with the processed input.
// 4. Formats the results (type U) into a markdown table.
// 5. Prints the formatted output table.
//
// This function uses reflection to handle various input and output types,
// making it flexible for different use cases. It can adapt to both individual
// items and slices/arrays, depending on the function signature provided.
//
// If any errors occur during parsing, processing, or function execution,
// Run will print an error message and return without producing output.
func Run[T, U any](f func(T) (U, error), inputTable string, opts ...RunOption) {
	c := RunConfig{}

	for _, opt := range opts {
		opt(&c)
	}

	table := parseMarkdown(inputTable)

	header := table[0]
	body := table[1:]

	var us []U

	switch reflect.TypeFor[T]().Kind() {
	case reflect.Slice, reflect.Array:
		// func f expects a slice of items as T, pass whole table to func
		t := reflect.New(reflect.TypeFor[T]()).Elem()
		for i, row := range body {
			elem := reflect.New(t.Type().Elem()).Elem()
			err := parseRow(elem, header, row)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			switch t.Kind() {
			case reflect.Slice:
				t = reflect.Append(t, elem)
			case reflect.Array:
				t.Index(i).Set(elem)
			}
		}
		u, err := f(t.Interface().(T))
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		us = append(us, u)
	default:
		// func f expects a single item, iterate and append for each item
		for _, row := range body {
			t := reflect.New(reflect.TypeFor[T]()).Elem()
			err := parseRow(t, header, row)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			u, err := f(t.Interface().(T))
			if err != nil {
				fmt.Printf("error: %v\n", err)
				return
			}
			us = append(us, u)
		}
	}

	outHeaders := c.outputHeaders
	if outHeaders == nil {
		t := reflect.TypeFor[U]()
		if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
			t = t.Elem()
		}
		for i := range t.NumField() {
			outHeaders = append(outHeaders, t.Field(i).Name)
		}
	}

	outTable := [][]string{outHeaders}

	for _, u := range us {
		v := reflect.ValueOf(u)
		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			for i := range v.Len() {
				rv := v.Index(i)
				var row []string
				for _, header := range outHeaders {
					f := rv.FieldByName(header)
					if !f.IsValid() {
						fmt.Printf("error: %v\n", fmt.Errorf("unknown field: %s", header))
						return
					}
					row = append(row, fmt.Sprint(f.Interface()))
				}
				outTable = append(outTable, row)
			}
		default:
			var row []string
			for _, header := range outHeaders {
				f := v.FieldByName(header)
				if !f.IsValid() {
					fmt.Printf("error: %v\n", fmt.Errorf("unknown field: %s", header))
					return
				}
				row = append(row, fmt.Sprint(f.Interface()))
			}
			outTable = append(outTable, row)
		}
	}

	revAlias := invert(c.headerAliases)

	for i, h := range outTable[0] {
		if a, ok := revAlias[h]; ok {
			outTable[0][i] = a
		}
	}

	fmt.Println(formatMarkdown(outTable))
}

func invert[T, U comparable](m map[T]U) map[U]T {
	m2 := make(map[U]T, len(m))
	for k, v := range m {
		m2[v] = k
	}
	return m2
}

// RunConfig holds configuration options for the Run function.
type RunConfig struct {
	outputHeaders []string
	headerAliases map[string]string
}

// RunOption is a function type used to set options for the Run function.
type RunOption func(*RunConfig)

// WithOutputHeader sets the output headers using a markdown-formatted header row.
func WithOutputHeader(header string) RunOption {
	return func(config *RunConfig) {
		config.outputHeaders = parseMarkdown(header)[0]
	}
}

// WithOutputHeaders sets the output headers using a slice of strings.
func WithOutputHeaders(headers []string) RunOption {
	return func(config *RunConfig) {
		config.outputHeaders = headers
	}
}

// WithHeaderAliases is a RunOption that sets up aliases for field names.
// This can be useful when the table headers don't exactly match
// the struct field names, allowing for more flexible parsing.
// In the aliases map, each key is an alias and its corresponding
// value is the struct field name it represents.
func WithHeaderAliases(aliases map[string]string) RunOption {
	return func(config *RunConfig) {
		config.headerAliases = aliases
	}
}

func parseRow(v reflect.Value, headers []string, row []string) error {
	for i, h := range headers {
		field := v.FieldByName(h)
		switch field.Kind() {
		case reflect.String:
			field.SetString(row[i])
		case reflect.Float64:
			var v float64
			_, err := fmt.Sscanf(row[i], "%f", &v)
			if err != nil {
				return err
			}
			field.SetFloat(v)
		default:
			return fmt.Errorf("unsupported type: %v\n", field.Kind())
		}
	}
	return nil
}
