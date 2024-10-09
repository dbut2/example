package example

import (
	"fmt"
	"regexp"
	"strings"
)

// parseMarkdown takes a Markdown table string and returns a table of strings
func parseMarkdown(m string) [][]string {
	m = strings.TrimSpace(m)
	rows := strings.Split(m, "\n")

	hasHeader := true
	if len(rows) < 2 {
		hasHeader = false
	} else {
		for _, char := range rows[1] {
			if char != '|' && char != '-' {
				hasHeader = false
				break
			}
		}
	}

	var header []string
	if hasHeader {
		header = splitRow(rows[0])
		rows = rows[2:]
	}

	var table [][]string
	if hasHeader {
		table = append(table, header)
	}

	for _, row := range rows {
		table = append(table, splitRow(row))
	}

	return table
}

func splitRow(row string) []string {
	re := regexp.MustCompile(`(^|[^\\])(\|)`)
	split := re.Split(row, -1)
	for i, cell := range split {
		cell = strings.ReplaceAll(cell, `\|`, `|`)
		cell = strings.TrimSpace(cell)
		split[i] = cell
	}
	return split[1 : len(split)-1]
}

func formatMarkdown(table [][]string) string {
	colWidth := make([]int, len(table[0]))
	for _, row := range table {
		for i, cell := range row {
			colWidth[i] = max(colWidth[i], len(cell))
		}
	}

	header := table[0]
	body := table[1:]

	output := new(strings.Builder)
	for i, cell := range header {
		fmt.Fprint(output, "|")
		fmt.Fprintf(output, " %-*s ", colWidth[i], cell)
		if i == len(header)-1 {
			fmt.Fprintln(output, "|")
		}
	}

	for i := range header {
		fmt.Fprint(output, "|")
		fmt.Fprint(output, strings.Repeat("-", colWidth[i]+2))
		if i == len(header)-1 {
			fmt.Fprintln(output, "|")
		}
	}

	for _, row := range body {
		for i, cell := range row {
			fmt.Fprint(output, "|")
			fmt.Fprintf(output, " %-*s ", colWidth[i], cell)
			if i == len(row)-1 {
				fmt.Fprintln(output, "|")
			}
		}
	}

	return output.String()
}
