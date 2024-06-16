package compiler

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"fmt"
	"reflect"
	"testing"
)

func TestCompiler_Compile(t *testing.T) {
	t.Run("returns error when lexing fails", func(t *testing.T) {
		lexerError := fmt.Errorf("fake lexing error")
		sourceCode := "fake source code"
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return nil, lexerError
			},
		}
		compilerArgs := Args{
			SourceCode: sourceCode,
		}
		expectedError := fmt.Errorf("error lexing program: %v", lexerError)

		compiler := NewCompiler(lexer, nil, nil, noOpPrint, noOpPrintProgramTree)
		_, err := compiler.Compile(compilerArgs)

		if err.Error() != expectedError.Error() {
			t.Errorf("Expected %v, got %v", lexerError, err)
		}
	})

	t.Run("returns error when parsing fails", func(t *testing.T) {
		parserError := fmt.Errorf("fake parsing error")
		sourceCode := "fake source code"
		tokens := []domain.Token{{Type: domain.Number}}
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return tokens, nil
			},
		}
		parser := &testutils.FakeParser{
			MockParse: func(tokens []domain.Token) (*domain.ProgramTree, error) {
				return nil, parserError
			},
		}
		compilerArgs := Args{
			SourceCode: sourceCode,
		}
		expectedError := fmt.Errorf("error parsing program: %v", parserError)

		compiler := NewCompiler(lexer, parser, nil, noOpPrint, noOpPrintProgramTree)
		_, err := compiler.Compile(compilerArgs)

		if err.Error() != expectedError.Error() {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when emitting fails", func(t *testing.T) {
		emitterError := fmt.Errorf("fake emitting error")
		sourceCode := "fake source code"
		tokens := []domain.Token{{Type: domain.Number}}
		programTree := &domain.ProgramTree{}
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return tokens, nil
			},
		}
		parser := &testutils.FakeParser{
			MockParse: func(tokens []domain.Token) (*domain.ProgramTree, error) {
				return programTree, nil
			},
		}
		emitter := &testutils.FakeEmitter{
			MockEmit: func(programTree *domain.ProgramTree) (string, error) {
				return "", emitterError
			},
		}
		compilerArgs := Args{
			SourceCode: sourceCode,
		}
		expectedError := fmt.Errorf("error emitting program: %v", emitterError)

		compiler := NewCompiler(lexer, parser, emitter, noOpPrint, noOpPrintProgramTree)
		_, err := compiler.Compile(compilerArgs)

		if err.Error() != expectedError.Error() {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}

	})

	t.Run("compiles program", func(t *testing.T) {
		sourceCode := "fake source code"
		tokens := []domain.Token{{Type: domain.Number}}
		programTree := &domain.ProgramTree{}
		compiledCode := "fake compiled code"
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return tokens, nil
			},
		}
		parser := &testutils.FakeParser{
			MockParse: func(tokens []domain.Token) (*domain.ProgramTree, error) {
				return programTree, nil
			},
		}
		emitter := &testutils.FakeEmitter{
			MockEmit: func(programTree *domain.ProgramTree) (string, error) {
				return compiledCode, nil
			},
		}
		compilerArgs := Args{
			SourceCode: sourceCode,
		}

		compiler := NewCompiler(lexer, parser, emitter, noOpPrint, noOpPrintProgramTree)
		actualCompiledCode, err := compiler.Compile(compilerArgs)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actualCompiledCode != compiledCode {
			t.Errorf("Expected %s, got %s", compiledCode, actualCompiledCode)
		}
	})

	t.Run("compiles program and show logs WHEN ShouldShowLogs is true", func(t *testing.T) {
		sourceCode := "fake source code"
		tokens := []domain.Token{{Type: domain.Number}}
		programTree := &domain.ProgramTree{}
		compiledCode := "fake compiled code"
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return tokens, nil
			},
		}
		parser := &testutils.FakeParser{
			MockParse: func(tokens []domain.Token) (*domain.ProgramTree, error) {
				return programTree, nil
			},
		}
		emitter := &testutils.FakeEmitter{
			MockEmit: func(programTree *domain.ProgramTree) (string, error) {
				return compiledCode, nil
			},
		}
		compilerArgs := Args{
			SourceCode:            sourceCode,
			ShouldShowLogs:        true,
			ShouldShowProgramTree: false,
		}
		var actualPrints []string
		printfMock := func(format string, v ...interface{}) (n int, err error) {
			actualPrints = append(actualPrints, fmt.Sprintf(format, v...))
			return 0, nil
		}
		expectedPrints := []string{
			"Lexing program\n",
			"Parsing program\n",
			"Emitting program\n",
		}

		compiler := NewCompiler(lexer, parser, emitter, printfMock, noOpPrintProgramTree)
		actualCompiledCode, err := compiler.Compile(compilerArgs)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actualCompiledCode != compiledCode {
			t.Errorf("Expected %s, got %s", compiledCode, actualCompiledCode)
		}
		if !reflect.DeepEqual(actualPrints, expectedPrints) {
			t.Errorf("Expected %v, got %v", expectedPrints, actualPrints)
		}
	})

	t.Run("compiles program and show logs and program tree WHEN ShouldShowLogs is true AND ShouldShowProgramTree is true", func(t *testing.T) {
		sourceCode := "fake source code"
		tokens := []domain.Token{{Type: domain.Number}}
		programTree := &domain.ProgramTree{}
		compiledCode := "fake compiled code"
		lexer := &testutils.FakeLexer{
			MockLex: func(inputCode string) ([]domain.Token, error) {
				return tokens, nil
			},
		}
		parser := &testutils.FakeParser{
			MockParse: func(tokens []domain.Token) (*domain.ProgramTree, error) {
				return programTree, nil
			},
		}
		emitter := &testutils.FakeEmitter{
			MockEmit: func(programTree *domain.ProgramTree) (string, error) {
				return compiledCode, nil
			},
		}
		compilerArgs := Args{
			SourceCode:            sourceCode,
			ShouldShowLogs:        true,
			ShouldShowProgramTree: true,
		}
		var actualPrints []string
		printfMock := func(format string, v ...interface{}) (n int, err error) {
			actualPrints = append(actualPrints, fmt.Sprintf(format, v...))
			return 0, nil
		}
		printProgramTreeMock := func(tree *domain.ProgramTree) {
			actualPrints = append(actualPrints, "Printed program tree")
		}
		expectedPrints := []string{
			"Lexing program\n",
			"Parsing program\n",
			"Program tree:\n",
			"Printed program tree",
			"Emitting program\n",
		}

		compiler := NewCompiler(lexer, parser, emitter, printfMock, printProgramTreeMock)
		actualCompiledCode, err := compiler.Compile(compilerArgs)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actualCompiledCode != compiledCode {
			t.Errorf("Expected %s, got %s", compiledCode, actualCompiledCode)
		}
		if !reflect.DeepEqual(actualPrints, expectedPrints) {
			t.Errorf("Expected %v, got %v", expectedPrints, actualPrints)
		}
	})
}

func noOpPrint(format string, v ...interface{}) (n int, err error) {
	panic("printf run but it should not have been called")
}

func noOpPrintProgramTree(tree *domain.ProgramTree) {
	panic("printProgramTree run but it should not have been called")
}
