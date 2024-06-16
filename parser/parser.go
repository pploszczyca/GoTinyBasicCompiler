package parser

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
)

type Parser interface {
	Parse(tokens []domain.Token) (*domain.ProgramTree, error)
}

type NodeParser interface {
	Parse(iterator *domain.TokenIterator) (*domain.Node, error)
}

type parser struct {
	lineParser NodeParser
}

func NewParser(
	lineParser NodeParser,
) Parser {
	return &parser{
		lineParser: lineParser,
	}
}

func (p *parser) Parse(tokens []domain.Token) (*domain.ProgramTree, error) {
	iterator := domain.NewTokenIterator(tokens)
	programTree := &domain.ProgramTree{}
	lineIndex := 1

	for iterator.HasNext() {
		node, err := p.lineParser.Parse(&iterator)
		if err != nil {
			return nil, fmt.Errorf("error parsing line %d: %v", lineIndex, err)
		}
		programTree.Nodes = append(programTree.Nodes, node)
		lineIndex++
	}

	return programTree, nil
}
