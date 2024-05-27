package parser

import "GoTinyBasicCompiler/domain"

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

func parseLine(tokens []domain.Token, currentIndex int) (*domain.ProgramTree, error) {

	return nil, nil
}
