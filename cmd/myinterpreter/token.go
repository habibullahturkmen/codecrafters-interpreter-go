package main

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

var reservedWords = map[string]string{
	"and":    "AND",
	"class":  "CLASS",
	"else":   "ELSE",
	"false":  "FALSE",
	"for":    "FOR",
	"fun":    "FUN",
	"if":     "IF",
	"nil":    "NIL",
	"or":     "OR",
	"print":  "PRINT",
	"return": "RETURN",
	"super":  "SUPER",
	"this":   "THIS",
	"true":   "TRUE",
	"var":    "VAR",
	"while":  "WHILE",
}
