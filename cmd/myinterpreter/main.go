package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
)

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
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	for _, item := range fileContents {
		switch item {
		case '(':
			fmt.Println(LEFT_PAREN, string(item), "null")
		case ')':
			fmt.Println(RIGHT_PAREN, string(item), "null")
		}
	}
	fmt.Println("EOF  null")
}
