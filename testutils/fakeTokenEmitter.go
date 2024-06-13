package testutils

import "GoTinyBasicCompiler/domain"

type FakeTokenEmitter struct {
	EmitMock func(token domain.Token) (string, error)
}

func (f FakeTokenEmitter) Emit(token domain.Token) (string, error) {
	return f.EmitMock(token)
}
