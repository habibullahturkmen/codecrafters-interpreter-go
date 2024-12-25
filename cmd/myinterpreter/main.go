package main

import (
	"fmt"
	"os"
)

type Token struct {
	Key   string
	Value rune
}

const (
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
)

var tokens = map[string]rune{
	"LEFT_PAREN":  '(',
	"RIGHT_PAREN": ')',
	"LEFT_BRACE":  '{',
	"RIGHT_BRACE": '}',
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	rawFileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContents := []rune(string(rawFileContents))
	for _, item := range fileContents {
		switch item {
		case tokens[LEFT_PAREN]:
			fmt.Printf("%s %c %s\n", LEFT_PAREN, item, "null")
		case tokens[RIGHT_PAREN]:
			fmt.Printf("%s %c %s\n", RIGHT_PAREN, item, "null")
		case tokens[LEFT_BRACE]:
			fmt.Printf("%s %c %s\n", LEFT_BRACE, item, "null")
		case tokens[RIGHT_BRACE]:
			fmt.Printf("%s %c %s\n", RIGHT_BRACE, item, "null")
		}
	}
	fmt.Println("EOF  null")
}
