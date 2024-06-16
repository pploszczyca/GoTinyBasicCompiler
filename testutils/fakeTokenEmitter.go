package testutils

import "GoTinyBasicCompiler/domain"

type FakeTokenEmitter struct {
	EmitMock func(token domain.Token) (string, error)
}

func (f FakeTokenEmitter) Emit(token domain.Token) (string, error) {
	if f.EmitMock == nil {
		panic("TokenEmitter.Emit run but it should not have been called")
	}

	return f.EmitMock(token)
}
