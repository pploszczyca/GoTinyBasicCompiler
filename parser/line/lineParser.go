package line

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type lineParser struct {
	statementParser parser.NodeParser
}

func NewLineParser(
	statementParser parser.NodeParser,
) parser.NodeParser {
	return &lineParser{
		statementParser: statementParser,
	}
}

func (l lineParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	lineNode := domain.Node{Type: domain.LineNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	if token.Type == domain.Number {
		lineNode.AddChild(&domain.Node{Type: domain.NumberNode, Token: token})
		iterator.Next()
	}

	statementNode, err := l.statementParser.Parse(iterator)
	if err != nil {
		return nil, err
	}

	lineNode.AddChild(statementNode)

	token, err = iterator.Current()
	if err != nil {
		return nil, err
	}

	if token.Type != domain.Cr {
		return nil, fmt.Errorf("expected CR token, but got %v", token.Type)
	}

	iterator.Next()

	return &lineNode, nil
}
