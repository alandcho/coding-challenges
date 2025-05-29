package parser

import (
	"fmt"
	"json_parser/model"
	"json_parser/model/ast"
	"strconv"
)

type ParserInterface interface {
	Parse(tokens *model.TokenList) (ast.ASTNode, error)
}

type Parser struct {

}

func New() ParserInterface {
	return &Parser{}
}


func (s *Parser) Parse(tokens *model.TokenList) (ast.ASTNode, error) {
	if (len(tokens.Tokens) == 0) {
		return ast.ASTNode{}, fmt.Errorf("empty token array")
	}
	
	return s.ParseValue(tokens)
}

func (s *Parser) ParseValue(tokens *model.TokenList) (ast.ASTNode, error) {
	token := tokens.Tokens[tokens.Index]
	switch token.TokenType {
		case model.TokenString:
			return ast.ASTNode{NodeType: ast.String, Value: token.Value}, nil
		case model.Number: 
			number, err := strconv.ParseFloat(token.Value, 64); 
			if err != nil {
				return ast.ASTNode{}, err
			}
			return ast.ASTNode{NodeType: ast.Number, Value: number}, nil
		case model.True:
			return ast.ASTNode{NodeType: ast.Boolean, Value: true}, nil
		case model.False:
			return ast.ASTNode{NodeType: ast.Boolean, Value: false}, nil
		case model.Null:
			return ast.ASTNode{NodeType: ast.Null, Value: nil}, nil
		case model.BraceOpen:
			return s.parseObject(tokens)
		case model.BracketOpen:
			return s.parseArray(tokens)
		default:
			return ast.ASTNode{}, fmt.Errorf("unexpected token type: %v", token.TokenType)
	} 
}


func (s *Parser) parseObject(tokens *model.TokenList) (ast.ASTNode, error) {
	node := ast.ASTNode{NodeType: ast.Object, Value: make(map[string]interface{})}
	token := tokens.NextToken()

	for token.TokenType != model.BraceClose {
		if (token.TokenType == model.TokenString) {
			key := token.Value
			token = tokens.NextToken()
			if (token.TokenType != model.Colon) {
				return ast.ASTNode{}, fmt.Errorf("expected : in key value pair")
			}
			token = tokens.NextToken()
			value, err := s.ParseValue(tokens)

			if err != nil {
				return ast.ASTNode{}, err
			}
			valueMap := node.Value.(map[string]interface{})
			valueMap[key] = value
			node.Value = valueMap
		}

		token = tokens.NextToken()
		if token.TokenType == model.Comma {
			token = tokens.NextToken()
		}
	}

	return node, nil
}


func (s *Parser) parseArray(tokens *model.TokenList) (ast.ASTNode, error) {
	node := ast.ASTNode{NodeType: ast.Array, Value: make([]interface{}, 0, 10)}
	token := tokens.NextToken() // Eat [

	for token.TokenType != model.BracketClose {
		value, err := s.ParseValue(tokens)

		if err != nil {
			return ast.ASTNode{}, err
		}
		valueSlice := node.Value.([]interface{})
		valueSlice = append(valueSlice, value)
		node.Value = valueSlice
		
		token = tokens.NextToken()
		if token.TokenType == model.Comma {
			token = tokens.NextToken()
		}
	}

	

	return ast.ASTNode{}, nil
}

