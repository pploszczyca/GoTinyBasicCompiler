package varList

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type varListParser struct {
}

func NewVarListParser() parser.NodeParser {
	return &varListParser{}
}

func (v varListParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	varListNode := domain.Node{Type: domain.VarListNode}

	for {
		token, err := iterator.Current()
		if err != nil {
			return nil, err
		}
		if token.Type != domain.Identifier {
			return nil, fmt.Errorf("expected identifier")
		}
		varListNode.AddChildToken(token)
		iterator.Next()

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
