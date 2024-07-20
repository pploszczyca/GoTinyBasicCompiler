package varList

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
)

type varListParser struct {
}

func NewVarListParser() parser.NodeParser {
	return &varListParser{}
}

func (v varListParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	varListNode := domain.Node{Type: domain.VarListNode}

	for {
		if err := internal.ExpectAndAddMatchingToken(iterator, &varListNode, domain.Identifier); err != nil {
			return nil, err
		}

		commaToken, err := iterator.Current()
		if err != nil {
			return nil, err
		}
		if commaToken.Type != domain.Comma {
			break
		}
		varListNode.AddChildToken(commaToken)
		iterator.Next()
	}

	return &varListNode, nil
}
