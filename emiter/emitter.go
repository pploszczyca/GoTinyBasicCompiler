package emiter

import (
	"GoTinyBasicCompiler/domain"
	"strings"
)

type Emitter interface {
	Emit(programTree *domain.ProgramTree) (string, error)
}

type cEmitter struct {
	tokenEmitter TokenEmitter
}

func NewCEmitter(
	tokenEmitter TokenEmitter,
) Emitter {
	return &cEmitter{
		tokenEmitter: tokenEmitter,
	}
}

func (c *cEmitter) Emit(programTree *domain.ProgramTree) (string, error) {
	var builder strings.Builder

	builder.WriteString("#include <stdio.h>\n")
	builder.WriteString("int main() {\n")

	for _, node := range programTree.Nodes {
		err := c.emitNode(&builder, node, 1)
		if err != nil {
			return "", err
		}
	}

	builder.WriteString("}\n")

	return builder.String(), nil
}

func (c *cEmitter) emitNode(builder *strings.Builder, node *domain.Node, indent int) error {
	if node.IsLeaf() {
		stringToken, err := c.tokenEmitter.Emit(node.Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken)
		return nil
	}

	switch node.Type {
	case domain.LineNode:
		return c.emitLineNode(builder, node, indent)
	case domain.ProgramNode:
	case domain.StatementNode:
		return c.emitStatementNode(builder, node, indent)
	case domain.ExpressionNode:
		for _, child := range node.Children {
			err := c.emitNode(builder, child, indent)
			if err != nil {
				return err
			}
		}
	case domain.NumberNode:
	case domain.IdentifierNode:
	case domain.OperatorNode:
	case domain.ExpressionListNode:
		for _, child := range node.Children {
			err := c.emitNode(builder, child, indent)
			if err != nil {
				return err
			}
		}
	case domain.TermNode:
		for _, child := range node.Children {
			err := c.emitNode(builder, child, indent)
			if err != nil {
				return err
			}
		}
	case domain.FactorNode:
		for _, child := range node.Children {
			err := c.emitNode(builder, child, indent)
			if err != nil {
				return err
			}
		}
	case domain.RelopNode:
	case domain.VarListNode:
	}

	return nil
}

func (c *cEmitter) emitLineNode(builder *strings.Builder, node *domain.Node, indent int) error {
	statementIndex := 0
	builder.WriteString(strings.Repeat("    ", indent))
	if node.Children[0].Type == domain.NumberNode {
		builder.WriteString("label_" + node.Children[0].Token.Value + ":\n")
		statementIndex = 1
		builder.WriteString(strings.Repeat("    ", indent))
	}

	for _, child := range node.Children[statementIndex:] {
		err := c.emitNode(builder, child, indent)
		if err != nil {
			return err
		}
	}
	builder.WriteString(";\n")
	return nil
}

func (c *cEmitter) emitStatementNode(builder *strings.Builder, node *domain.Node, indent int) error {
	switch node.Children[0].Token.Type {
	case domain.Print:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken)
		builder.WriteString("(\"")
		expressionListNode := node.Children[1]

		for _, child := range expressionListNode.Children {
			if child.Token.Type == domain.String {
				builder.WriteString("%s")
			} else if child.Token.Type == domain.Comma {
				builder.WriteString(",")
			} else {
				builder.WriteString("%d")
			}
		}

		builder.WriteString("\", ")

		err = c.emitNode(builder, expressionListNode, indent)
		if err != nil {
			return err
		}

		builder.WriteString(")")
	}

	return nil
}
