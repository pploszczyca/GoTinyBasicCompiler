package parser

import "GoTinyBasicCompiler/domain"

type expressionListParser struct {
}

func NewExpressionListParser() NodeParser {
	return &expressionListParser{}
}

func (e expressionListParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	expressionListNode := domain.Node{Type: domain.ExpressionNode}

	if tokens[currentIndex].Type == domain.String {
		expressionListNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	}

	// TODO: Implement parsing expression and list of expressions

	return &expressionListNode, currentIndex, nil
}
