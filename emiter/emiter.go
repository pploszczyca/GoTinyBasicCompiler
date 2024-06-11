package emiter

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"strings"
)

type ProgramCode string

type Emitter interface {
	Emit(programTree domain.ProgramTree) (ProgramCode, error)
}

type cEmitter struct {
}

func NewCEmitter() Emitter {
	return &cEmitter{}
}

func (c *cEmitter) Emit(programTree domain.ProgramTree) (ProgramCode, error) {
	var builder strings.Builder

	builder.WriteString("#include <stdio.h>\n")
	builder.WriteString("int main() {\n")

	for _, node := range programTree.Nodes {
		c.emitNode(&builder, node, 1)
	}

	builder.WriteString("  return 0;\n")
	builder.WriteString("}\n")

	return ProgramCode(builder.String()), nil
}

func (c *cEmitter) emitNode(builder *strings.Builder, node *domain.Node, indent int) {
	prefix := strings.Repeat("  ", indent)
	switch node.Type {
	case domain.LineNode:
		for _, child := range node.Children {
			c.emitNode(builder, child, indent)
		}
	case domain.StatementNode:
		for _, child := range node.Children {
			c.emitNode(builder, child, indent)
		}
	case domain.ExpressionNode:
		builder.WriteString(prefix)
		c.emitExpression(builder, node)
		builder.WriteString(";\n")
	case domain.ExpressionListNode:
		builder.WriteString(prefix)
		c.emitExpression(builder, node)
		builder.WriteString(";\n")
	case domain.TermNode:
		c.emitTerm(builder, node)
	case domain.FactorNode:
		c.emitFactor(builder, node)
	case domain.RelopNode:
		c.emitRelop(builder, node)
	case domain.VarListNode:
		c.emitVarList(builder, node)
	default:
		fmt.Printf("Unknown node type: %v\n", node.Type)
	}
}

func (c *cEmitter) emitExpression(builder *strings.Builder, node *domain.Node) {
	if len(node.Children) > 0 {
		for i, child := range node.Children {
			if i > 0 {
				builder.WriteString(" ")
			}
			c.emitNode(builder, child, 0)
		}
	} else {
		c.emitToken(builder, node.Token)
	}
}

func (c *cEmitter) emitTerm(builder *strings.Builder, node *domain.Node) {
	for _, child := range node.Children {
		c.emitNode(builder, child, 0)
	}
}

func (c *cEmitter) emitFactor(builder *strings.Builder, node *domain.Node) {
	if len(node.Children) > 0 {
		for _, child := range node.Children {
			c.emitNode(builder, child, 0)
		}
	} else {
		c.emitToken(builder, node.Token)
	}
}

func (c *cEmitter) emitRelop(builder *strings.Builder, node *domain.Node) {
	c.emitToken(builder, node.Token)
}

func (c *cEmitter) emitVarList(builder *strings.Builder, node *domain.Node) {
	for i, child := range node.Children {
		if i > 0 {
			builder.WriteString(", ")
		}
		c.emitNode(builder, child, 0)
	}
}

func (c *cEmitter) emitToken(builder *strings.Builder, token domain.Token) {
	switch token.Type {
	case domain.Number:
		builder.WriteString(token.Value)
	case domain.Identifier:
		builder.WriteString(token.Value)
	case domain.String:
		builder.WriteString(fmt.Sprintf("\"%s\"", token.Value))
	case domain.Print:
		builder.WriteString("printf")
	case domain.Input:
		builder.WriteString("scanf")
	case domain.Plus:
		builder.WriteString("+")
	case domain.Minus:
		builder.WriteString("-")
	case domain.Multiply:
		builder.WriteString("*")
	case domain.Divide:
		builder.WriteString("/")
	case domain.Equal:
		builder.WriteString("=")
	case domain.LessThan:
		builder.WriteString("<")
	case domain.GreaterThan:
		builder.WriteString(">")
	case domain.LessThanOrEqual:
		builder.WriteString("<=")
	case domain.GreaterThanOrEqual:
		builder.WriteString(">=")
	case domain.NotEqual:
		builder.WriteString("!=")
	case domain.Comma:
		builder.WriteString(", ")
	case domain.Semicolon:
		builder.WriteString(";")
	case domain.LParen:
		builder.WriteString("(")
	case domain.RParen:
		builder.WriteString(")")
	default:
		fmt.Printf("Unknown token type: %v\n", token.Type)
	}
}
