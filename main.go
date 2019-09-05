package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"os"
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
			//fmt.Printf("found at [%d] %s\n", i, currentLine)
			return i
		}
	}
	return -1
}

func findMethodLine(lines []string, idx int) string {
	for i := idx - 1; i >= 0; i-- {
		prevLine := lines[i]
		if strings.HasSuffix(prevLine, "(") {
			//fmt.Printf("found at [%d] %s\n", i, prevLine)
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

func readBytesFromFile() []byte {
	targetSrcFile := os.Args[1]
	content, err := ioutil.ReadFile(targetSrcFile)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func readBytesFromPipe() []byte {
	reader := bufio.NewReader(os.Stdin)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func readLines() []string {
	var content []byte
	if len(os.Args) > 1 {
		content = readBytesFromFile()
	} else {
		content = readBytesFromPipe()
	}
	//fmt.Printf("File contents: %s", content)
	lines := strings.Split(string(content), "\n")
	return lines
}

func parse(lines []string) []string {
	//) {
	for idx, line := range lines {
		if strings.HasSuffix(line, ") {") && !strings.Contains(line, "(") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			//fmt.Printf("[%s]\n", spaces)
			lines[idx] = line[:len(line)-3] + "\n" + spaces + ") {"
		}
	}
	//fmt.Printf("%s", strings.Join(lines, "\n"))

	//) throws ** {
	for idx, line := range lines {
		if strings.Contains(line, ") throws ") && !strings.Contains(line, "(") && strings.HasSuffix(line, "{") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			//fmt.Printf("[%s]\n", spaces)
			splitPoint := strings.Index(line, ") throws ")
			lines[idx] = line[:splitPoint] + "\n" + spaces + line[splitPoint:]
		}
	}
	//fmt.Printf("%s", strings.Join(lines, "\n"))

	//) throws MalformedURLException,
	//         IllegalStateException {
	for idx, line := range lines {
		if strings.Contains(line, ") throws ") && !strings.Contains(line, "(") && strings.HasSuffix(line, ",") {
			prevLine := findMethodLine(lines, idx)
			spaces := extractSpaces(prevLine)
			//fmt.Printf("[%s]\n", spaces)
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
	//fmt.Printf("%v\n", lines)
	fmt.Printf("%s", strings.Join(parse(lines), "\n"))
}
