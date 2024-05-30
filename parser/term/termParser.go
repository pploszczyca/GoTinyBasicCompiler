package term

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
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

	factorNode, err := t.factorParser.Parse(iterator)
	if err != nil {
		return nil, err
	}
	termNode.AddChild(factorNode)

	for {
		token, err := iterator.Current()
		if err != nil {
			return nil, err
		}

		if token.Type == domain.Multiply || token.Type == domain.Divide {
			termNode.AddChild(&domain.Node{Token: token})
			iterator.Next()

			factorNode, err := t.factorParser.Parse(iterator)
			if err != nil {
				return nil, err
			}
			termNode.AddChild(factorNode)
		} else {
			break
		}
	}

	return termNode, nil
}
