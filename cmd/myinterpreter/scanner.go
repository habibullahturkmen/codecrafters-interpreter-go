package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func scanner(fileContents []rune) {
	line := 1
	errors := 0
	contentLength := len(fileContents)
	if contentLength > 0 {
		for i := 0; i < contentLength; i++ {
			token := fileContents[i]
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
				if i+1 < contentLength && fileContents[i+1] == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", EqualEqual, token, token, "null")
					i++ // Skip the second "=".
				} else {
					fmt.Printf("%s %c %s\n", EQUAL, token, "null")
				}
			case tokens[BANG]: // Handle "!" and "!=".
				if i+1 < contentLength && fileContents[i+1] == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", BangEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", BANG, token, "null")
				}
			case tokens[LESS]: // Handle "<" and "<=".
				if i+1 < contentLength && fileContents[i+1] == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", LessEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", LESS, token, "null")
				}
			case tokens[GREATER]: // Handle ">" and ">=".
				if i+1 < contentLength && fileContents[i+1] == tokens[EQUAL] {
					fmt.Printf("%s %c%c %s\n", GreaterEqual, token, tokens[EQUAL], "null")
					i++ // Skip the "=".
				} else {
					fmt.Printf("%s %c %s\n", GREATER, token, "null")
				}
			case tokens[SLASH]: // Handle "/" and "//" (comments).
				if i+1 < contentLength && fileContents[i+1] == tokens[SLASH] {
					// Skip single-line comments and increment line count for each newline encountered.
					for i < contentLength && fileContents[i] != tokens[NEWLINE] {
						i++
					}
					line++
				} else {
					fmt.Printf("%s %c %s\n", SLASH, token, "null")
				}
			case tokens[STRING]:
				start := i
				for i+1 < contentLength && fileContents[i+1] != tokens[STRING] {
					if fileContents[i+1] == tokens[NEWLINE] {
						line++
					}
					i++
				}
				if i+1 >= contentLength {
					errors++
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
				} else {
					i++
					stringLexeme := string(fileContents[start : i+1])
					stringLiteral := string(fileContents[start+1 : i])
					fmt.Printf("%s %s %s\n", STRING, stringLexeme, stringLiteral)
				}
			case tokens[NEWLINE]:
				line++
			case tokens[TAB], tokens[SPACE]:
				continue
			default:
				if unicode.IsDigit(token) {
					start := i
					for i+1 < contentLength && (unicode.IsDigit(fileContents[i+1]) || fileContents[i+1] == '.') {
						i++
					}
					numberLexeme := string(fileContents[start : i+1])
					fmt.Printf("%s %s %s\n", NUMBER, numberLexeme, parseNumberLiteral(numberLexeme))
				} else if isAlpha(token) {
					start := i
					for i+1 < contentLength && isAlphaNumeric(fileContents[i+1]) {
						i++
					}
					identifier := string(fileContents[start : i+1])
					if len(reservedWords[identifier]) != 0 {
						fmt.Printf("%s %s %s\n", reservedWords[identifier], identifier, "null")
					} else {
						fmt.Printf("%s %s %s\n", IDENTIFIER, identifier, "null")
					}
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
