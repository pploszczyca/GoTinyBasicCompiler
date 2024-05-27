package parser

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
)

type Parser interface {
	Parse(tokens []domain.Token) (domain.ProgramTree, error)
}

type parser struct {
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) Parse(tokens []domain.Token) (domain.ProgramTree, error) {
	return domain.ProgramTree{}, nil
}

func parseLine(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	lineNode := domain.Node{Type: domain.LineNode}

	if tokens[currentIndex].Type == domain.Number {
		lineNode.AddChild(&domain.Node{Type: domain.NumberNode, Token: tokens[currentIndex]})
		currentIndex++
	}

	statementNode, currentIndex, err := parseStatement(tokens, currentIndex)
	if err != nil {
		return nil, currentIndex, err
	}

	lineNode.AddChild(statementNode)

	if tokens[currentIndex].Type != domain.Cr {
		return nil, currentIndex, fmt.Errorf("expected CR token, but got %v", tokens[currentIndex].Type)
	}

	currentIndex++

	return &lineNode, currentIndex, nil
}

func parseStatement(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	statementNode := domain.Node{Type: domain.StatementNode}

	switch tokens[currentIndex].Type {
	case domain.Print:
		statementNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
		expressionListNode, currentIndex, err := parseExpressionList(tokens, currentIndex)
		if err != nil {
			return nil, currentIndex, err
		}
		statementNode.AddChild(expressionListNode)

	default:
		return nil, currentIndex, fmt.Errorf("unexpected statement: %v", tokens[currentIndex].Type)
	}

	return nil, -1, nil
}

func parseExpressionList(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	expressionListNode := domain.Node{Type: domain.ExpressionNode}

	if tokens[currentIndex].Type == domain.String {
		expressionListNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
	}

	// TODO: Implement parsing expression and list of expressions

	return &expressionListNode, currentIndex, nil
}
