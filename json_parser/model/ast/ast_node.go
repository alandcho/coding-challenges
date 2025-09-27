package ast

import "fmt"

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

func (a ASTNode) String() string {
	return fmt.Sprintf("ASTNode{Type: %v, Value: %v}", a.NodeType, a.Value)
}