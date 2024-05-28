package domain

type ProgramTree struct {
	Nodes []*Node
}

type Node struct {
	Type     NodeType
	Token    Token
	Children []*Node
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
)
