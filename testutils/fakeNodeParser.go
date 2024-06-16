package testutils

import "GoTinyBasicCompiler/domain"

type FakeNodeParser struct {
	ParseMock func(iterator *domain.TokenIterator) (*domain.Node, error)
}

func (f FakeNodeParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	if f.ParseMock == nil {
		panic("NodeParser.Parse run but it should not have been called")
	}
	return f.ParseMock(iterator)
}
