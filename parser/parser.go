package parser

import (
	"GoTinyBasicCompiler/domain"
)

type Parser interface {
	Parse(tokens []domain.Token) (domain.ProgramTree, error)
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

func (p *parser) Parse(tokens []domain.Token) (domain.ProgramTree, error) {
	iterator := domain.NewTokenIterator(tokens)
	programTree := domain.ProgramTree{}

	for iterator.HasNext() {
		node, err := p.lineParser.Parse(iterator)
		if err != nil {
			return domain.ProgramTree{}, err
		}
		programTree.Nodes = append(programTree.Nodes, node)
	}

	return programTree, nil
}
