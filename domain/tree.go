package domain

import "fmt"

type ProgramTree struct {
	Nodes []*Node
}

type Node struct {
	Type     NodeType
	Token    Token
	Children []*Node
}

func (n *Node) AddChildToken(token Token) {
	n.AddChild(&Node{Token: token})
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) IsLeaf() bool {
	return len(n.Children) == 0 && n.Token != Token{}
}

type NodeType int

const (
	ProgramNode NodeType = iota
	StatementNode
	ExpressionNode
	NumberNode
	IdentifierNode
	OperatorNode
	LineNode
	ExpressionListNode
	TermNode
	FactorNode
	RelopNode
	VarListNode
)

var nodeTypeToString = map[NodeType]string{
	ProgramNode:        "ProgramNode",
	StatementNode:      "StatementNode",
	ExpressionNode:     "ExpressionNode",
	NumberNode:         "NumberNode",
	IdentifierNode:     "IdentifierNode",
	OperatorNode:       "OperatorNode",
	LineNode:           "LineNode",
	ExpressionListNode: "ExpressionListNode",
	TermNode:           "TermNode",
	FactorNode:         "FactorNode",
	RelopNode:          "RelopNode",
	VarListNode:        "VarListNode",
}

func (n NodeType) String() string {
	if str, ok := nodeTypeToString[n]; ok {
		return str
	}
	return fmt.Sprintf("Unknown NodeType (%d)", n)
}
