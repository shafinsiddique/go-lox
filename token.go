package main

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
