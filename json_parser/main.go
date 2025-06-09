package main

import (
	"fmt"
	"json_parser/internal/parser"
	"json_parser/internal/tokenizer"
	"json_parser/model"
)


func main() {
	jsonString := `{
  "id": "647ceaf3657eade56f8224eb",
  "index": 0,
  "anArray": [],
  "boolean": true,
  "nullValue": null,
  "nestedObject": {
	"key": "value"
  }
}`
	tokenizerService := tokenizer.NewTokenizer()
	parserService := parser.New()
	
	token, err := tokenizerService.Tokenize(jsonString)
	if err != nil {
		fmt.Printf("Tokenization error: %v\n", err)
		return
	}

	tokenList := model.TokenList{Tokens: token, Index: 0}
	rootNode, err := parserService.Parse(&tokenList)
	if err != nil {
		fmt.Printf("Parsing error: %v\n", err)
		return
	}

	fmt.Println(rootNode)
}