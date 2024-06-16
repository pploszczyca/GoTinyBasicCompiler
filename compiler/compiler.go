package compiler

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/emiter"
	"GoTinyBasicCompiler/lexer"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type Args struct {
	SourceCode            string
	ShouldShowLogs        bool
	ShouldShowProgramTree bool
}

type Compiler interface {
	Compile(args Args) (string, error)
}

type compiler struct {
	lexer            lexer.Lexer
	parser           parser.Parser
	emitter          emiter.Emitter
	printf           func(format string, v ...any) (n int, err error)
	printProgramTree func(tree *domain.ProgramTree)
}

func NewCompiler(
	lexer lexer.Lexer,
	parser parser.Parser,
	emitter emiter.Emitter,
	printf func(format string, v ...any) (n int, err error),
	printProgramTree func(tree *domain.ProgramTree),
) Compiler {
	return &compiler{
		lexer:            lexer,
		parser:           parser,
		emitter:          emitter,
		printf:           printf,
		printProgramTree: printProgramTree,
	}
}

func (c *compiler) Compile(args Args) (string, error) {
	// TODO: Add time measurements
	c.printIfRequired("Lexing program", args.ShouldShowLogs)
	tokens, err := c.lexer.Lex(args.SourceCode)
	if err != nil {
		return "", fmt.Errorf("error lexing program: %v", err)
	}

	c.printIfRequired("Parsing program", args.ShouldShowLogs)
	programTree, err := c.parser.Parse(tokens)
	if err != nil {
		return "", fmt.Errorf("error parsing program: %v", err)
	}

	if args.ShouldShowProgramTree {
		_, _ = c.printf("Program tree:\n")
		c.printProgramTree(programTree)
	}

	c.printIfRequired("Emitting program", args.ShouldShowLogs)
	compiledCode, err := c.emitter.Emit(programTree)
	if err != nil {
		return "", fmt.Errorf("error emitting program: %v", err)
	}

	return compiledCode, nil
}

func (c *compiler) printIfRequired(message string, showLogs bool) {
	if showLogs {
		_, _ = c.printf(message)
	}
}
