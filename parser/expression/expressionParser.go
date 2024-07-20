package expression

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
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

		if err := internal.ParseAndAddNode(iterator, expressionNode, e.termParser); err != nil {
			return nil, err
		}

		token, err = iterator.Current()
		if err != nil {
			return nil, err
		}
		if token.Type != domain.Plus && token.Type != domain.Minus {
			break
		}
	}

	return expressionNode, nil
}
