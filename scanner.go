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
	if val, ok := OneCharLexemes[currentChar]; ok {
		*token = Token{Lexeme: currentChar, TokenType: val, Line: *line}
	} else if currentChar == "!" {
	}

}
