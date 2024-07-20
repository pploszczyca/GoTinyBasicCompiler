package factor

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
	"fmt"
)

type factorParser struct {
	expressionParser parser.NodeParser
}

func NewFactorParser(
	expressionParser parser.NodeParser,
) parser.NodeParser {
	return &factorParser{
		expressionParser: expressionParser,
	}
}

func (f factorParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	factorNode := &domain.Node{Type: domain.FactorNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	if token.Type == domain.Number || token.Type == domain.Identifier {
		factorNode.AddChildToken(token)
		iterator.Next()
	} else if token.Type == domain.LParen {
		factorNode.AddChildToken(token)
		iterator.Next()

		if err := internal.ParseSteps([]internal.StepFunc{
			func() error { return internal.ParseAndAddNode(iterator, factorNode, f.expressionParser) },
			func() error { return internal.ExpectAndAddMatchingToken(iterator, factorNode, domain.RParen) },
		}); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("expected number, identifier or left parenthesis")
	}

	return factorNode, nil
}
