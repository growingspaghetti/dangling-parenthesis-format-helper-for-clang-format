package main

import (
	"fmt"
	"strings"
)

func parse(lines []string) []string {
	//) {
	for idx, line := range lines {
		if strings.HasSuffix(line, ") {") && !strings.Contains(line, "(") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			lines[idx] = line[:len(line)-3] + "\n" + spaces + ") {"
		}
	}

	//) throws ** {
	for idx, line := range lines {
		if strings.Contains(line, ") throws ") && !strings.Contains(line, "(") && strings.HasSuffix(line, "{") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			splitPoint := strings.Index(line, ") throws ")
			lines[idx] = line[:splitPoint] + "\n" + spaces + line[splitPoint:]
		}
	}

	//) throws MalformedURLException,
	//         IllegalStateException {
	for idx, line := range lines {
		if strings.Contains(line, ") throws ") && !strings.Contains(line, "(") && strings.HasSuffix(line, ",") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			braketLineIdx := findNextOpenBracket(lines, idx)
			if braketLineIdx == -1 {
				continue
			}
			splitPoint := strings.Index(line, ") throws ")
			lines[idx] = line[:splitPoint] + "\n" + spaces + line[splitPoint:]

			reduceIndentation(lines, idx+1, braketLineIdx, spaces)

			braketLine := lines[braketLineIdx]
			lines[braketLineIdx] = braketLine[:len(braketLine)-2] + "\n" + spaces + "{"
		}
	}
	return lines
}

func main() {
	lines := readLines()
	fmt.Printf("%s", strings.Join(parse(lines), "\n"))
}
