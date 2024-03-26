package locknock

import (
	"html/template"
	"strings"
	"unicode"
)

var templateHelpers *template.FuncMap

func init() {
	templateHelpers = &template.FuncMap{
		"dedent": dedent,
		"sum":    sum,
	}
}

// dedent removes indentation and leading spaces so multiline strigs become more readable
func dedent(text string) string {
	lines := strings.Split(text, "\n")
	minIndent := -1
	for _, line := range lines {
		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		currentIndent := 0
		for _, rune := range line {
			if !unicode.IsSpace(rune) {
				break
			}
			currentIndent++
		}
		if minIndent == -1 || currentIndent < minIndent {
			minIndent = currentIndent
		}
	}
	if minIndent > 0 {
		for i, line := range lines {
			if len(line) >= minIndent {
				lines[i] = line[minIndent:]
			}
		}
	}
	// trim leading spaces
	leadingIdx := 0
	for leadingIdx = 0; leadingIdx < len(lines); leadingIdx++ {
		leadingString := lines[leadingIdx]
		leadingString = strings.TrimLeftFunc(leadingString, unicode.IsSpace)
		lines[leadingIdx] = leadingString
		if len(leadingString) > 0 {
			break
		}
	}
	lines = lines[leadingIdx:]
	return strings.Join(lines, "\n")
}

// sum is a custom function that sums to ints
func sum(a, b int) int {
	return a + b
}
