package main

var OneCharLexemes = map[string]TokenType{
	":": COLON,
	"(": LEFT_PAREN,
	")": RIGHT_PAREN,
	";": SEMI_COLON,
	"+": ADD,
	"-": SUBTRACT,
	"*": MULTIPY,
	"/": DIVIDE,
}
