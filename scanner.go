package main

import "fmt"

func GetTokensFromSource(source string) string {
	current := 0
	line := 1
	var tokens []Token
	for current < len(source) {
		scanToken(source, &tokens, &current, &line)
	}

	x := 0

	for x < len(tokens) {
		fmt.Println(tokens[x].ToString())
		x += 1
	}

	return "hello"
}

func scanToken(source string, tokens *[]Token, currentPosition *int, line *int) {
	token := Token{}
	currentChar := string(source[*currentPosition])

	if _, ok := TwoCharLexemes[currentChar]; ok {
		token = getTokenFromTwoCharLexeme(source, currentPosition, *line)
	} else if val, ok := AllLexemes[currentChar]; ok {
		token = Token{Lexeme: currentChar, TokenType: val, Line: *line}
		*currentPosition += 1
	} else if source[*currentPosition] == '\n' || source[*currentPosition] == '\r' {
		handleNewLineCharacter(source, currentPosition)
		*line += 1
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
		peeker := string(source[position+1])
		fmt.Println(peeker)
		if string(source[position+1]) == character {
			return true
		}
	}
	return false
}

func handleNewLineCharacter(source string, current *int) {
	if source[*current] == '\r' {
		*current += 2
	} else {
		*current += 1
	}
}
