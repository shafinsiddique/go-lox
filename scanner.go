package main

import "fmt"

func isLetter(character byte) bool {
	if character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' {
		return true
	}
	return false
}
func GetTokensFromSource(source string) []Token {
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

	return tokens
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
		handleNewLineCharacter(source, currentPosition, line)
	} else if source[*currentPosition] == '"' {
		token = handleStrings(source, currentPosition, line)
	} else if source[*currentPosition] == ' ' {
		*currentPosition += 1
	} else if isLetter(source[*currentPosition]) {
		token = handleReservedWord(source, currentPosition, line)
	}
	// to handle : numbers and reserved words.
	if token.IsInitialized() {
		*tokens = append(*tokens, token)
	}

}

func handleReservedWord(source string, currentPosition *int, line *int) Token {
	var word string

	for *currentPosition < len(source) && !isLetter(source[*currentPosition]) {
		word += string(source[*currentPosition])
		*currentPosition += 1
	}

}

func handleStrings(source string, currentPosition *int, line *int) Token {
	var str string = string(source[*currentPosition])
	startingLine := *line
	*currentPosition += 1

	for *currentPosition < len(source) && source[*currentPosition] != '"' {
		str += string(source[*currentPosition])

		if source[*currentPosition] == '\r' || source[*currentPosition] == '\n' {
			handleNewLineCharacter(source, currentPosition, line)
		} else {
			*currentPosition += 1
		}
	}

	if *currentPosition == len(source) {
		return Token{}

	} else {
		str += "\""
		*currentPosition += 1
		return Token{Lexeme: str, TokenType: STRING, Line: startingLine}
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

func handleNewLineCharacter(source string, current *int, line *int) {
	if source[*current] == '\r' {
		*current += 2
	} else {
		*current += 1
	}
	*line += 1
}
