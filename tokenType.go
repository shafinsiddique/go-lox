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
