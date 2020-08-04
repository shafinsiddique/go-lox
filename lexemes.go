package main

var AllLexemes = map[string]TokenType{
	":":  COLON,
	"(":  LEFT_PAREN,
	")":  RIGHT_PAREN,
	";":  SEMI_COLON,
	"+":  ADD,
	"-":  SUBTRACT,
	"*":  MULTIPY,
	"/":  DIVIDE,
	"!":  EXCLAMATION,
	"!=": NOT_EXCLAMATION,
	"=":  EQUAL,
	"==": EQUAL_EQUAL,
}

var TwoCharLexemes = map[string]string{
	"!": "=",
	"=": "=",
	">": "<",
	"<": "=",
}
