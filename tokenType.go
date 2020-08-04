package main

type TokenType struct {
	Name string
}

var LEFT_PAREN = TokenType{Name: "LEFT_PAREN"}

var RIGHT_PAREN = TokenType{Name: "RIGHT_PAREN"}

var SEMI_COLON = TokenType{Name: "SEMI_COLON"}

var COLON = TokenType{Name: "COLON"}

var ADD = TokenType{Name: "ADD"}

var SUBTRACT = TokenType{Name: "SUBTRACT"}

var MULTIPY = TokenType{Name: "MULTIPLY"}

var DIVIDE = TokenType{Name: "DIVIDE"}

var EQUAL_EQUAL = TokenType{Name: "EQUAL_EQUAL"}

var EQUAL = TokenType{Name: "EQUAL_EQUAL"}

var EXCLAMATION = TokenType{Name: "EXCLAMATION"}

var NOT_EXCLAMATION = TokenType{Name: "NOT_EXCLAMATION"}

var GREATER = TokenType{Name: "GREATER"}

var LESSER = TokenType{Name: "LESSER"}

var GREATER_EQUAL = TokenType{Name: "GREATER_EQUAL"}

var LESSER_EQUAL = TokenType{Name: "LESSER_EQUAL"}
