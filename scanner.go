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
	switch currentChar := source[current]; currentChar {
	case '(':
		*token = Token{Lexeme: "(", Line: *line, TokenType: LEFT_PAREN}
	case ')':
		*token = Token{Lexeme: "(", Line: *line, TokenType: RIGHT_PAREN}
	case '\n':
		*line += 1
	case ':':
		*token = Token{Lexeme: ":", Line: *line, TokenType: COLON}
	case ';':
		*token = Token{Lexeme: ";", Line: *line, TokenType: SEMI_COLON}
	case '+':
		*token = Token{Lexeme: "+", Line: *line, TokenType: ADD}

	case '-':
		*token = Token{Lexeme: "-", Line: *line, TokenType: SUBTRACT}

	case '*':
		*token = Token{Lexeme: "*", Line: *line, TokenType: MULTIPY}

	case '/':
		*token = Token{Lexeme: "/", Line: *line, TokenType: DIVIDE}
	}

	if token != nil {
		*tokens = append(*tokens, *token)

	}
}
