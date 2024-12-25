package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA       = "COMMA"
	DOT         = "DOT"
	MINUS       = "MINUS"
	PLUS        = "PLUS"
	SEMICOLON   = "SEMICOLON"
	STAR        = "STAR"
	SLASH       = "SLASH"
)

var tokens = map[string]rune{
	"LEFT_PAREN":  '(',
	"RIGHT_PAREN": ')',
	"LEFT_BRACE":  '{',
	"RIGHT_BRACE": '}',
	"COMMA":       ',',
	"DOT":         '.',
	"MINUS":       '-',
	"PLUS":        '+',
	"SEMICOLON":   ';',
	"STAR":        '*',
	"SLASH":       '/',
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
	if len(fileContents) > 0 {
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
			case tokens[COMMA]:
				fmt.Printf("%s %c %s\n", COMMA, item, "null")
			case tokens[DOT]:
				fmt.Printf("%s %c %s\n", DOT, item, "null")
			case tokens[MINUS]:
				fmt.Printf("%s %c %s\n", MINUS, item, "null")
			case tokens[PLUS]:
				fmt.Printf("%s %c %s\n", PLUS, item, "null")
			case tokens[SEMICOLON]:
				fmt.Printf("%s %c %s\n", SEMICOLON, item, "null")
			case tokens[STAR]:
				fmt.Printf("%s %c %s\n", STAR, item, "null")
			case tokens[SLASH]:
				fmt.Printf("%s %c %s\n", SLASH, item, "null")
			}
		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null")
	}
}
