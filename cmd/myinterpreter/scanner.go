package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var parsedTokens []string

func scanner(fileContents []rune) {
	line := 1
	errors := 0
	contentLength := len(fileContents)

	if contentLength > 0 {
		for i := 0; i < contentLength; i++ {
			currentToken := fileContents[i]
			switch currentToken {
			case '(':
				addToken(token{tokenType: LeftParen})
			case ')':
				addToken(token{tokenType: RightParen})
			case '{':
				addToken(token{tokenType: LeftBrace})
			case '}':
				addToken(token{tokenType: RightBrace})
			case ',':
				addToken(token{tokenType: COMMA})
			case '.':
				addToken(token{tokenType: DOT})
			case '-':
				addToken(token{tokenType: MINUS})
			case '+':
				addToken(token{tokenType: PLUS})
			case ';':
				addToken(token{tokenType: SEMICOLON})
			case '*':
				addToken(token{tokenType: STAR})
			case '=': // Handle "=" and "==".
				if i+1 < contentLength && fileContents[i+1] == '=' {
					addToken(token{tokenType: EqualEqual})
					i++ // Skip the second "=".
				} else {
					addToken(token{tokenType: EQUAL})
				}
			case '!': // Handle "!" and "!=".
				if i+1 < contentLength && fileContents[i+1] == '=' {
					addToken(token{tokenType: BangEqual})
					i++ // Skip the "=".
				} else {
					addToken(token{tokenType: BANG})
				}
			case '<': // Handle "<" and "<=".
				if i+1 < contentLength && fileContents[i+1] == '=' {
					addToken(token{tokenType: LessEqual})
					i++ // Skip the "=".
				} else {
					addToken(token{tokenType: LESS})
				}
			case '>': // Handle ">" and ">=".
				if i+1 < contentLength && fileContents[i+1] == '=' {
					addToken(token{tokenType: GreaterEqual})
					i++ // Skip the "=".
				} else {
					addToken(token{tokenType: GREATER})
				}
			case '/': // Handle "/" and "//" (comments).
				if i+1 < contentLength && fileContents[i+1] == '/' {
					// Skip single-line comments and increment line count for each newline encountered.
					for i < contentLength && fileContents[i] != '\n' {
						i++
					}
					line++
				} else {
					addToken(token{tokenType: SLASH})
				}
			case '"':
				start := i
				for i+1 < contentLength && fileContents[i+1] != '"' {
					if fileContents[i+1] == '\n' {
						line++
					}
					i++
				}
				if i+1 >= contentLength {
					errors++
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
				} else {
					i++
					lexeme := string(fileContents[start : i+1])
					literal := string(fileContents[start+1 : i])
					addToken(token{tokenType: STRING, lexeme: lexeme, literal: literal})
				}
			case '\n':
				line++
			case '\t', ' ', '\r':
				continue
			default:
				if unicode.IsDigit(currentToken) {
					start := i
					for i+1 < contentLength && (unicode.IsDigit(fileContents[i+1]) || fileContents[i+1] == '.') {
						i++
					}
					lexeme := string(fileContents[start : i+1])
					addToken(token{tokenType: NUMBER, lexeme: lexeme, literal: parseNumberLiteral(lexeme)})
				} else if isAlpha(currentToken) {
					start := i
					for i+1 < contentLength && isAlphaNumeric(fileContents[i+1]) {
						i++
					}
					lexeme := string(fileContents[start : i+1])
					if len(reservedWords[lexeme]) != 0 {
						addToken(token{tokenType: reservedWords[lexeme], lexeme: lexeme, literal: "null"})
					} else {
						addToken(token{tokenType: IDENTIFIER, lexeme: lexeme, literal: "null"})
					}

				} else {
					errors++
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, currentToken)
				}
			}
		}
		addToken(token{tokenType: EOF})
	} else {
		addToken(token{tokenType: EOF})
	}

	for _, token := range parsedTokens {
		fmt.Print(token)
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

func addToken(token token) {
	if token.lexeme == "" && token.literal == "" {
		parsedTokens = formatAndAppendToken(token.tokenType, tokens[token.tokenType], "null")
	} else {
		parsedTokens = formatAndAppendToken(token.tokenType, token.lexeme, token.literal)
	}
}

func formatAndAppendToken(tokenType string, lexeme string, literal string) []string {
	return append(parsedTokens, fmt.Sprintf("%s %s %s\n", tokenType, lexeme, literal))
}
