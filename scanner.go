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
	token := new(Token)
	currentChar := string(source[current])

	if _, ok := TwoCharLexemes[currentChar]; ok {
		handleTwoCharLexeme(source, currentPosition, *line)
	} else if val, ok := AllLexemes[currentChar]; ok {
		*token = Token{Lexeme: currentChar, TokenType: val, Line: *line}
	}

}

func handleTwoCharLexeme(source string, currentPosition *int, line int) {
	currentChar := string(source[*currentPosition])
	token := new(Token)
	val, _ := TwoCharLexemes[currentChar]
	lexeme := currentChar
	if peek(source, *currentPosition, val) {
		*currentPosition += 2
		lexeme += val
	}
	tokenType, _ := AllLexemes[lexeme]
	*token = Token{Lexeme: lexeme, TokenType: tokenType, Line: line}
}

func peek(source string, position int, character string) bool {
	if position < len(source)-1 {
		if string(source[position+1]) == character {
			return true
		}
	}
	return false
}
