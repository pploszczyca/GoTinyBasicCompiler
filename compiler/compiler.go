package compiler

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/emiter"
	"GoTinyBasicCompiler/lexer"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/utils"
	"fmt"
	"time"
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
	printf           func(format string, v ...any)
	printProgramTree func(tree *domain.ProgramTree)
}

func NewCompiler(
	lexer lexer.Lexer,
	parser parser.Parser,
	emitter emiter.Emitter,
	printf func(format string, v ...any),
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
	c.printIfRequired("Lexing program\n", args.ShouldShowLogs)
	tokens, duration, err := utils.MeasureTime(func() ([]domain.Token, error) {
		return c.lexer.Lex(args.SourceCode)
	})
	c.printDurationIfRequired("Lexing", duration, args.ShouldShowLogs)
	if err != nil {
		return "", fmt.Errorf("error lexing program: %v", err)
	}

	c.printIfRequired("Parsing program\n", args.ShouldShowLogs)
	programTree, duration, err := utils.MeasureTime(func() (*domain.ProgramTree, error) {
		return c.parser.Parse(tokens)
	})
	c.printDurationIfRequired("Parsing", duration, args.ShouldShowLogs)
	if err != nil {
		return "", fmt.Errorf("error parsing program: %v", err)
	}

	if args.ShouldShowProgramTree {
		c.printf("Program tree:\n")
		c.printProgramTree(programTree)
	}

	c.printIfRequired("Emitting program\n", args.ShouldShowLogs)
	compiledCode, duration, err := utils.MeasureTime(func() (string, error) {
		return c.emitter.Emit(programTree)
	})
	c.printDurationIfRequired("Emitting", duration, args.ShouldShowLogs)
	if err != nil {
		return "", fmt.Errorf("error emitting program: %v", err)
	}

	return compiledCode, nil
}

func (c *compiler) printIfRequired(message string, showLogs bool) {
	if showLogs {
		c.printf(message)
	}
}

func (c *compiler) printDurationIfRequired(testCase string, duration time.Duration, showLogs bool) {
	if showLogs {
		c.printf("%s time elapsed: %.6f seconds\n", testCase, duration.Seconds())
	}
}
