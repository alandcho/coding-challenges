package tokenizer

import (
	"fmt"
	"json_parser/model"
	"strconv"
	"strings"
	"unicode"
)

type TokenizerInterface interface {
	Tokenize(input string) ([]model.Token, error)
}

type Tokenizer struct {

}

func NewTokenizer() TokenizerInterface {
	return &Tokenizer{}
}

var (
	charToTokenTypeMap = map[byte]model.TokenType{
		'{': model.BraceOpen,
		'}': model.BraceClose,
		'[': model.BracketOpen,
		']': model.BracketClose,
		',': model.Comma,
		':': model.Colon,
	}
	stringToTokenTypeMap = map[string]model.TokenType{
		"true": model.True,
		"false": model.False,
		"null": model.Null,
	} 
)

func (s *Tokenizer) Tokenize(input string) ([]model.Token, error) {
	index := 0
	result := []model.Token{}

	for index < len(input) {
		char := input[index]

		tokenType, exist := charToTokenTypeMap[char]
		if exist {
			result = append(result, model.Token{TokenType: tokenType, Value: string(char)})
		} else if char == '"' {
			// Process string
			var sb strings.Builder
			index++
			char = input[index]
			for char != '"' {
				sb.WriteByte(char)
				index++
				char = input[index]
			}
			result = append(result, model.Token{TokenType: model.TokenString, Value: sb.String()})
		} else if unicode.IsDigit(rune(char)) || unicode.IsLetter(rune(char)) {
			// Process other value as string
			var sb strings.Builder
			for unicode.IsDigit(rune(char)) || unicode.IsLetter(rune(char)) {
				sb.WriteByte(char)
				index++
				char = input[index]
			}
			processedString := sb.String()
			tokenType, exist := stringToTokenTypeMap[processedString]

			if exist {
				result = append(result, model.Token{TokenType: tokenType, Value: processedString})
			} else if _, err := strconv.ParseFloat(processedString, 64); err == nil {
				result = append(result, model.Token{TokenType: model.Number, Value: processedString})
			} else {
				return nil, fmt.Errorf("failed when parsing this value 1: %v", processedString)
			}
		} else if unicode.IsSpace(rune(char)) {
			index++
			continue
		} else {
			// If reach this then it's unrecognized character
			return nil, fmt.Errorf("failed when parsing this value 2: %v", char)
		}

		index++
	}

	return result, nil
}

