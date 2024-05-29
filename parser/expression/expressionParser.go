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

func (e expressionParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	expressionNode := &domain.Node{Type: domain.ExpressionNode}

	if tokens[currentIndex].Type == domain.Plus || tokens[currentIndex].Type == domain.Minus {
		expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	}

	newIndex, node, i, err := e.parseFactor(tokens, currentIndex, expressionNode)
	if err != nil {
		return node, i, err
	}
	currentIndex = newIndex

	for {
		if tokens[currentIndex].Type == domain.Multiply || tokens[currentIndex].Type == domain.Divide {
			expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
			currentIndex++

			newIndex, node, i, err = e.parseFactor(tokens, currentIndex, expressionNode)
			if err != nil {
				return node, i, err
			}
			currentIndex = newIndex
			expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		} else {
			newIndex, node, i, err = e.parseFactor(tokens, currentIndex, expressionNode)
			if err != nil {
				break
			}
			currentIndex = newIndex
			expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		}
	}

	// TODO: Add parsing list of expressions

	return expressionNode, currentIndex, nil
}

func (e expressionParser) parseFactor(tokens []domain.Token, currentIndex int, expressionNode *domain.Node) (int, *domain.Node, int, error) {
	if tokens[currentIndex].Type == domain.Number || tokens[currentIndex].Type == domain.Identifier {
		expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	} else if tokens[currentIndex].Type == domain.LParen {
		expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
		expressionNode, newIndex, err := e.Parse(tokens, currentIndex)
		if err != nil {
			return 0, nil, newIndex, err
		}
		currentIndex = newIndex
		if tokens[currentIndex].Type != domain.RParen {
			return 0, nil, currentIndex, fmt.Errorf("expected RParen token, but got %v", tokens[currentIndex].Type)
		}
		expressionNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	} else {
		return 0, nil, currentIndex, fmt.Errorf("unexpected token: %v", tokens[currentIndex].Type)
	}
	return currentIndex, nil, 0, nil
}
