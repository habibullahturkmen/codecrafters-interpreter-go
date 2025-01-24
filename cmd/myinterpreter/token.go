package main

// Example output
// <TOKEN_TYPE>    <LEXEME>    <LITERAL>
// SLASH           /           null
// STRING          "foo baz"   foo baz

type token struct {
	tokenType string
	lexeme    string
	literal   string
}

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
	STRING       = "STRING"
	NUMBER       = "NUMBER"
	IDENTIFIER   = "IDENTIFIER"
	EOF          = "EOF"
)

var tokens = map[string]string{
	"LEFT_PAREN":    "(",
	"RIGHT_PAREN":   ")",
	"LEFT_BRACE":    "{",
	"RIGHT_BRACE":   "}",
	"COMMA":         ",",
	"DOT":           ".",
	"MINUS":         "-",
	"PLUS":          "+",
	"SEMICOLON":     ";",
	"STAR":          "*",
	"SLASH":         "/",
	"EQUAL":         "=",
	"EQUAL_EQUAL":   "==",
	"BANG":          "!",
	"BANG_EQUAL":    "!=",
	"LESS":          "<",
	"LESS_EQUAL":    "<=",
	"GREATER":       ">",
	"GREATER_EQUAL": ">=",
	"NEWLINE":       "\n",
	"TAB":           "\t",
	"SPACE":         " ",
	"STRING":        "\"",
	"EOF":           "",
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
