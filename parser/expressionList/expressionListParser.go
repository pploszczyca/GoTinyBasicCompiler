package expressionList

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
)

type expressionListParser struct {
	expressionParser parser.NodeParser
}

func NewExpressionListParser(
	expressionParser parser.NodeParser,
) parser.NodeParser {
	return &expressionListParser{
		expressionParser: expressionParser,
	}
}

func (e expressionListParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	expressionListNode := &domain.Node{Type: domain.ExpressionListNode}

	if tokens[currentIndex].Type == domain.String {
		expressionListNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	} else {
		expressionNode, newIndex, err := e.expressionParser.Parse(tokens, currentIndex)
		if err != nil {
			return nil, newIndex, err
		}
		currentIndex = newIndex
		expressionListNode.AddChild(expressionNode)
	}

	// TODO: Implement parsing list of expressions or strings

	return expressionListNode, currentIndex, nil
}
