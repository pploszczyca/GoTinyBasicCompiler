package line

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/statement"
	"fmt"
)

type lineParser struct {
	statementParser parser.NodeParser
}

func NewLineParser() parser.NodeParser {
	return &lineParser{
		statementParser: statement.NewStatementParser(),
	}
}

func (l lineParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	lineNode := domain.Node{Type: domain.LineNode}

	if tokens[currentIndex].Type == domain.Number {
		lineNode.AddChild(&domain.Node{Type: domain.NumberNode, Token: tokens[currentIndex]})
		currentIndex++
	}

	statementNode, currentIndex, err := l.statementParser.Parse(tokens, currentIndex)
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
