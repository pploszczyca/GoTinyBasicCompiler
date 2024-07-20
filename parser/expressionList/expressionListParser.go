package expressionList

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
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

		if err := internal.ExpectAndAddMatchingToken(iterator, expressionListNode, domain.Comma); err != nil {
			break
		}
	}

	return expressionListNode, nil
}
