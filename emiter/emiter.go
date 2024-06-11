package emiter

import "GoTinyBasicCompiler/domain"

type ProgramCode string

type Emitter interface {
	Emit(programTree domain.ProgramTree) (ProgramCode, error)
}
