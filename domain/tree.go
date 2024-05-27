package domain

type ProgramTree struct {
	Nodes []*Node
}

type Node struct {
	Type     NodeType
	Token    Token
	Children []*Node
}

type NodeType int

const (
	ProgramNode NodeType = iota
	StatementNode
	ExpressionNode
	NumberNode
	IdentifierNode
	OperatorNode
)
