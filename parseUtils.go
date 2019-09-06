package main

import (
	"bytes"
	"strings"
)

func reduceIndentation(lines []string, from int, to int, spaces string) {
	for i := from; i <= to; i++ {
		currentLine := lines[i]
		lines[i] = spaces + "         " + strings.TrimSpace(currentLine)
	}
}

func findNextOpenBracket(lines []string, idx int) int {
	for i := idx; i < len(lines); i++ {
		currentLine := lines[i]
		if strings.HasSuffix(currentLine, " {") {

			return i
		}
	}
	return -1
}

func findMethodLine(lines []string, idx int) string {
	for i := idx - 1; i >= 0; i-- {
		prevLine := lines[i]
		if strings.HasSuffix(prevLine, "(") {
			return prevLine
		}
	}
	return ""
}

func extractSpaces(line string) string {
	var buffer bytes.Buffer
	runes := []rune(line)
	for _, r := range runes {
		if r == ' ' {
			buffer.WriteString(" ")
		} else {
			break
		}
	}
	return buffer.String()
}
