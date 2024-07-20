package term

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
)

type termParser struct {
	factorParser parser.NodeParser
}

func NewTermParser(
	factorParser parser.NodeParser,
) parser.NodeParser {
	return &termParser{
		factorParser: factorParser,
	}
}

func (t termParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	termNode := &domain.Node{Type: domain.TermNode}

	for {
		if err := internal.ParseAndAddNode(iterator, termNode, t.factorParser); err != nil {
			return nil, err
		}

		token, err := iterator.Current()
		if err != nil {
			return nil, err
		}

		if token.Type != domain.Multiply && token.Type != domain.Divide {
			break
		}

		termNode.AddChildToken(token)
		iterator.Next()
	}

	return termNode, nil
}
