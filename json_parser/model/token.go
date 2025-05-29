package model


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

func (tk *TokenList) NextToken() Token{
	tk.Index = tk.Index + 1
	return tk.Tokens[tk.Index]
}