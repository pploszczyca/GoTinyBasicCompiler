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

func (e expressionListParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	expressionListNode := &domain.Node{Type: domain.ExpressionListNode}

	for {
		token, err := iterator.Current()
		if err != nil {
			return nil, err
		}

		if token.Type == domain.String {
			expressionListNode.AddChildToken(token)
			iterator.Next()
		} else {
			expressionNode, err := e.expressionParser.Parse(iterator)
			if err != nil {
				return nil, err
			}
			expressionListNode.AddChild(expressionNode)
		}

		token, err = iterator.Current()
		if err != nil {
			return nil, err
		}

		if token.Type != domain.Comma {
			break
		}
		expressionListNode.AddChildToken(token)
		iterator.Next()
	}

	return expressionListNode, nil
}
