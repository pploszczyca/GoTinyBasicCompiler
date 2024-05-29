package expression

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type expressionParser struct {
}

func NewExpressionParser() parser.NodeParser {
	return &expressionParser{}
}

func (e expressionParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	expressionNode := &domain.Node{Type: domain.ExpressionNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	if token.Type == domain.Plus || token.Type == domain.Minus {
		expressionNode.AddChild(&domain.Node{Token: token})
		iterator.Next()
	}

	err = e.parseFactor(iterator, expressionNode)
	if err != nil {
		return nil, err
	}

	token, err = iterator.Current()
	if err != nil {
		return nil, err
	}

	for {
		if token.Type == domain.Multiply || token.Type == domain.Divide {
			expressionNode.AddChild(&domain.Node{Token: token})
			iterator.Next()

			err = e.parseFactor(iterator, expressionNode)
			if err != nil {
				return nil, err
			}
		} else {
			err = e.parseFactor(iterator, expressionNode)
			if err != nil {
				break
			}
		}
	}

	// TODO: Add parsing list of expressions

	return expressionNode, nil
}

func (e expressionParser) parseFactor(iterator *domain.TokenIterator, expressionNode *domain.Node) error {
	token, err := iterator.Current()
	if err != nil {
		return err
	}

	if token.Type == domain.Number || token.Type == domain.Identifier {
		expressionNode.AddChild(&domain.Node{Token: token})
		iterator.Next()
	} else if token.Type == domain.LParen {
		expressionNode.AddChild(&domain.Node{Token: token})
		iterator.Next()
		expressionNode, err := e.Parse(iterator)
		if err != nil {
			return err
		}
		token, err = iterator.Current()
		if err != nil {
			return err
		}
		if token.Type != domain.RParen {
			return fmt.Errorf("expected RParen token, but got %v", token.Type)
		}
		expressionNode.AddChild(&domain.Node{Token: token})
		iterator.Next()
	} else {
		return fmt.Errorf("unexpected token: %v", token.Type)
	}
	return nil
}
