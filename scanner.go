package main

func GetTokensFromSource(source string) string {
	current := 0
	line := 1
	var tokens []Token
	for current < len(source) {
		scanToken(source, &tokens, &current, &line)
		//scanToken()
	}
	return "hello"
}

func scanToken(source string, tokens *[]Token, currentPosition *int, line *int) {
	current := *currentPosition
	token := Token{}
	currentChar := string(source[current])

	if _, ok := TwoCharLexemes[currentChar]; ok {
		token = getTokenFromTwoCharLexeme(source, currentPosition, *line)
	} else if val, ok := AllLexemes[currentChar]; ok {
		token = Token{Lexeme: currentChar, TokenType: val, Line: *line}
	}

	if token.IsInitialized() {
		*tokens = append(*tokens, token)
	}

}

func getTokenFromTwoCharLexeme(source string, currentPosition *int, line int) Token {
	currentChar := string(source[*currentPosition])
	val, _ := TwoCharLexemes[currentChar]
	lexeme := currentChar
	if peek(source, *currentPosition, val) {
		*currentPosition += 2
		lexeme += val
	}
	tokenType, _ := AllLexemes[lexeme]
	return Token{Lexeme: lexeme, TokenType: tokenType, Line: line}
}

func peek(source string, position int, character string) bool {
	if position < len(source)-1 {
		if string(source[position+1]) == character {
			return true
		}
	}
	return false
}
