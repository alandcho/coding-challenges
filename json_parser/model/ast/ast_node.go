package ast


type ASTNodeType int

const (
	Object ASTNodeType = iota
	Array
	String
	Number
	Boolean
	Null
)

type ASTNode struct {
	NodeType ASTNodeType
	Value interface{}
}