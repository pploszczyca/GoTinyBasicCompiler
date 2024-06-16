package testutils

import "GoTinyBasicCompiler/domain"

type FakeLexer struct {
	MockLex func(inputCode string) ([]domain.Token, error)
}

func (f *FakeLexer) Lex(inputCode string) ([]domain.Token, error) {
	if f.MockLex == nil {
		panic("Lexer.Lex run but it should not have been called")
	}
	return f.MockLex(inputCode)
}

type FakeParser struct {
	MockParse func(tokens []domain.Token) (*domain.ProgramTree, error)
}

func (f *FakeParser) Parse(tokens []domain.Token) (*domain.ProgramTree, error) {
	if f.MockParse == nil {
		panic("Parser.Parse run but it should not have been called")
	}
	return f.MockParse(tokens)
}

type FakeEmitter struct {
	MockEmit func(programTree *domain.ProgramTree) (string, error)
}

func (f *FakeEmitter) Emit(programTree *domain.ProgramTree) (string, error) {
	if f.MockEmit == nil {
		panic("Emitter.Emit run but it should not have been called")
	}
	return f.MockEmit(programTree)
}

type FakeNodeParser struct {
	ParseMock func(iterator *domain.TokenIterator) (*domain.Node, error)
}

func (f FakeNodeParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	if f.ParseMock == nil {
		panic("NodeParser.Parse run but it should not have been called")
	}
	return f.ParseMock(iterator)
}

type FakeTokenEmitter struct {
	EmitMock func(token domain.Token) (string, error)
}

func (f FakeTokenEmitter) Emit(token domain.Token) (string, error) {
	if f.EmitMock == nil {
		panic("TokenEmitter.Emit run but it should not have been called")
	}

	return f.EmitMock(token)
}
