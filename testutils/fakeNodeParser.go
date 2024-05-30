package testutils

import "GoTinyBasicCompiler/domain"

type FakeNodeParser struct {
	ParseMock func(iterator *domain.TokenIterator) (*domain.Node, error)
}

func (f FakeNodeParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	return f.ParseMock(iterator)
}
