package expression

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
)

type expressionParser struct {
}

func NewExpressionParser() parser.NodeParser {
	return &expressionParser{}
}

func (e expressionParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	expressionNode := &domain.Node{Type: domain.ExpressionNode}

	if tokens[currentIndex].Type == domain.Plus || tokens[currentIndex].Type == domain.Minus {
		expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	}

	// TODO: Add parsing of other expression types

	return expressionNode, currentIndex, nil
}
