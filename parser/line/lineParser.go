package line

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
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

	if err := internal.ParseAndAddNode(iterator, &lineNode, l.statementParser); err != nil {
		return nil, err
	}

	token, err = iterator.Current()
	if err != nil {
		return nil, err
	}

	if token.Type != domain.Cr && token.Type != domain.Eof {
		return nil, fmt.Errorf("expected CR token, but got %v", token.Type)
	}

	iterator.Next()

	return &lineNode, nil
}
