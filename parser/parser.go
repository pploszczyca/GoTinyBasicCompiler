package parser

import (
	"GoTinyBasicCompiler/domain"
)

type Parser interface {
	Parse(tokens []domain.Token) (domain.ProgramTree, error)
}

type NodeParser interface {
	Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error)
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
	currentIndex := 0
	programTree := domain.ProgramTree{}

	for currentIndex < len(tokens) {
		node, newIndex, err := p.lineParser.Parse(tokens, currentIndex)
		if err != nil {
			return domain.ProgramTree{}, err
		}
		currentIndex = newIndex
		programTree.Nodes = append(programTree.Nodes, node)
	}

	return programTree, nil
}
