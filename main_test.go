package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	actualc, _ := ioutil.ReadFile("a.java")
	actual := parse(strings.Split(string(actualc), "\n"))
	expectedc, _ := ioutil.ReadFile("expected.java")
	expected := strings.Split(string(expectedc), "\n")
	if strings.Join(actual, "\n") != strings.Join(expected, "\n") {
		t.Fatal("failed test")
	}
}
