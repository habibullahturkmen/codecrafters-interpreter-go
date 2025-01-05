package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	LeftParen    = "LEFT_PAREN"
	RightParen   = "RIGHT_PAREN"
	LeftBrace    = "LEFT_BRACE"
	RightBrace   = "RIGHT_BRACE"
	COMMA        = "COMMA"
	DOT          = "DOT"
	MINUS        = "MINUS"
	PLUS         = "PLUS"
	SEMICOLON    = "SEMICOLON"
	STAR         = "STAR"
	SLASH        = "SLASH"
	EQUAL        = "EQUAL"
	EqualEqual   = "EQUAL_EQUAL"
	BangEqual    = "BANG_EQUAL"
	BANG         = "BANG"
	LESS         = "LESS"
	LessEqual    = "LESS_EQUAL"
	GREATER      = "GREATER"
	GreaterEqual = "GREATER_EQUAL"
	NEWLINE      = "NEWLINE"
	TAB          = "TAB"
	SPACE        = "SPACE"
	STRING       = "STRING"
	NUMBER       = "NUMBER"
	IDENTIFIER   = "IDENTIFIER"
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
	"EQUAL":       '=',
	"BANG":        '!',
	"LESS":        '<',
	"GREATER":     '>',
	"NEWLINE":     '\n',
	"TAB":         '\t',
	"SPACE":       ' ',
	"STRING":      '"',
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

	// Replace "<|TAB|>" and "<|SPACE|>" with '\t' and ' ' if the file includes them
	var processedContents []rune
	for i := 0; i < len(fileContents); i++ {
		if i+6 < len(fileContents) && string(fileContents[i:i+7]) == "<|TAB|>" {
			processedContents = append(processedContents, '\t')
			i += 6 // Skip "<|TAB|>"
		} else if i+8 < len(fileContents) && string(fileContents[i:i+9]) == "<|SPACE|>" {
			processedContents = append(processedContents, ' ')
			i += 8 // Skip "<|SPACE|>"
		} else {
			processedContents = append(processedContents, fileContents[i])
		}
	}

	line := 1
	errors := 0
	contentLength := len(processedContents)
	if contentLength > 0 {
		for i := 0; i < contentLength; i++ {
			token := processedContents[i]
			nextToken := processedContents[i+1]
			switch token {
			case tokens[LeftParen]:
				fmt.Printf("%s %c %s\n", LeftParen, token, "null")
			case tokens[RightParen]:
				fmt.Printf("%s %c %s\n", RightParen, token, "null")
			case tokens[LeftBrace]:
				fmt.Printf("%s %c %s\n", LeftBrace, token, "null")
			case tokens[RightBrace]:
				fmt.Printf("%s %c %s\n", RightBrace, token, "null")
			case tokens[COMMA]:
				fmt.Printf("%s %c %s\n", COMMA, token, "null")
			case tokens[DOT]:
				fmt.Printf("%s %c %s\n", DOT, token, "null")
			case tokens[MINUS]:
				fmt.Printf("%s %c %s\n", MINUS, token, "null")
			case tokens[PLUS]:
				fmt.Printf("%s %c %s\n", PLUS, token, "null")
			case tokens[SEMICOLON]:
				fmt.Printf("%s %c %s\n", SEMICOLON, token, "null")
			case tokens[STAR]:
				fmt.Printf("%s %c %s\n", STAR, token, "null")
			case tokens[EQUAL]: // Handle "=" and "==".
				if i+1 < contentLength && nextToken == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", EqualEqual, token, token, "null")
					i++ // Skip the second "=".
				} else {
					fmt.Printf("%s %c %s\n", EQUAL, token, "null")
				}
			case tokens[BANG]: // Handle "!" and "!=".
				if i+1 < contentLength && nextToken == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", BangEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", BANG, token, "null")
				}
			case tokens[LESS]: // Handle "<" and "<=".
				if i+1 < contentLength && nextToken == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", LessEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", LESS, token, "null")
				}
			case tokens[GREATER]: // Handle ">" and ">=".
				if i+1 < contentLength && nextToken == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", GreaterEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", GREATER, token, "null")
				}
			case tokens[SLASH]: // Handle "/" and "//" (comments).
				if i+1 < contentLength && nextToken == tokens[SLASH] {
					// Skip single-line comments and increment line count for each newline encountered.
					for i < contentLength && processedContents[i] != tokens[NEWLINE] {
						i++
					}
					line++
				} else {
					fmt.Printf("%s %c %s\n", SLASH, token, "null")
				}
			case tokens[STRING]:
				start := i
				for i+1 < contentLength && nextToken != tokens[STRING] {
					if nextToken == tokens[NEWLINE] {
						line++
					}
					i++
				}
				if i+1 >= contentLength {
					errors++
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
				} else {
					i++
					stringLexeme := string(processedContents[start : i+1])
					stringLiteral := string(processedContents[start+1 : i])
					fmt.Printf("%s %s %s\n", STRING, stringLexeme, stringLiteral)
				}
			case tokens[NEWLINE]:
				line++
			case tokens[TAB], tokens[SPACE]:
				continue
			default:
				if unicode.IsDigit(token) {
					start := i
					for i+1 < contentLength && (unicode.IsDigit(nextToken) || nextToken == '.') {
						i++
					}
					numberLexeme := string(processedContents[start : i+1])
					fmt.Printf("%s %s %s\n", NUMBER, numberLexeme, parseNumberLiteral(numberLexeme))
				} else if isAlpha(token) {
					start := i
					for i+1 < contentLength && isAlpha(nextToken) {
						//if nextToken == tokens[NEWLINE] {
						//	line++
						//}
						i++
					}
					identifier := string(processedContents[start:i])
					fmt.Printf("%s %s %s\n", IDENTIFIER, identifier, "null")
				} else {
					errors++
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, token)
				}
			}
		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null")
	}

	if errors > 0 {
		os.Exit(65)
	}
}

func parseNumberLiteral(input string) string {
	floatValue, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return input
	}

	if floatValue == float64(int(floatValue)) {
		return fmt.Sprintf("%.1f", floatValue)
	}

	return strconv.FormatFloat(floatValue, 'f', -1, 64)
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || unicode.IsDigit(c)
}

// Example output
// <TOKEN_TYPE>    <LEXEME>    <LITERAL>
// SLASH           /           null
// STRING          "foo baz"   foo baz
