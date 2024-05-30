package expression

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
)

type expressionParser struct {
	termParser parser.NodeParser
}

func NewExpressionParser(
	termParser parser.NodeParser,
) parser.NodeParser {
	return &expressionParser{
		termParser: termParser,
	}
}

func (e expressionParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	expressionNode := &domain.Node{Type: domain.ExpressionNode}

	for {
		token, err := iterator.Current()
		if err != nil {
			return nil, err
		}

		if token.Type == domain.Plus || token.Type == domain.Minus {
			expressionNode.AddChildToken(token)
			iterator.Next()
		}

		termNode, err := e.termParser.Parse(iterator)
		if err != nil {
			return nil, err
		}
		expressionNode.AddChild(termNode)

		token, err = iterator.Current()
		if err != nil || (token.Type != domain.Plus && token.Type != domain.Minus) {
			break
		}
	}

	return expressionNode, nil
}
