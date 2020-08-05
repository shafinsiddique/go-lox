package main

import "strconv"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Line      int
}

func (token Token) IsInitialized() bool {
	if token.Lexeme != "" {
		return true
	}

	return false
}

func (token Token) ToString() string {
	return "Token Type: " + token.TokenType.Name + "\nLexeme: " + token.Lexeme + "\nLine: " + strconv.Itoa(token.Line)
}
