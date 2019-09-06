package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
