package relop

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type relopParser struct {
	allowedOperators map[domain.TokenType]bool
}

func NewRelopParser() parser.NodeParser {
	allowedOperators := map[domain.TokenType]bool{
		domain.LessThan:           true,
		domain.LessThanOrEqual:    true,
		domain.GreaterThan:        true,
		domain.GreaterThanOrEqual: true,
		domain.Equal:              true,
		domain.NotEqual:           true,
	}

	return &relopParser{
		allowedOperators: allowedOperators,
	}
}

func (r relopParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	relopNode := &domain.Node{Type: domain.RelopNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	if r.allowedOperators[token.Type] {
		relopNode.AddChildToken(token)
		iterator.Next()
	} else {
		return nil, fmt.Errorf("unexpected relop: %v", token.Type)
	}

	return relopNode, nil
}
