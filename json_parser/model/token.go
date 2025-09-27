package model

import "fmt"

type TokenType int

const (
	BraceOpen TokenType = iota
	BraceClose
	BracketOpen
	BracketClose
	TokenString
	Number
	Comma
	Colon
	True
	False
	Null	
)

type Token struct {
	TokenType TokenType
	Value string
}


type TokenList struct {
	Tokens []Token
	Index  int32
}

func (tl *TokenList) NextToken() Token{
	tl.Index = tl.Index + 1
	if int(tl.Index) >= len(tl.Tokens) {
		// Return a safe token instead of panicking
		return Token{TokenType: -1, Value: ""}
	}
	return tl.Tokens[tl.Index]
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %v, Value: %q}", t.TokenType, t.Value)
}