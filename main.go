package main

import (
	"GoTinyBasicCompiler/compiler"
	emiterModule "GoTinyBasicCompiler/emiter"
	lexerModule "GoTinyBasicCompiler/lexer"
	"GoTinyBasicCompiler/parser/expression"
	"GoTinyBasicCompiler/parser/expressionList"
	"GoTinyBasicCompiler/parser/factor"
	"GoTinyBasicCompiler/parser/line"
	"GoTinyBasicCompiler/parser/relop"
	"GoTinyBasicCompiler/parser/statement"
	"GoTinyBasicCompiler/parser/term"
	"GoTinyBasicCompiler/parser/varList"
	"GoTinyBasicCompiler/utils"
	"fmt"
	"log"
	"os"
)
import parserModule "GoTinyBasicCompiler/parser"

type ProgramArguments struct {
	ProgramPath string
	FilePath    string
	OutputPath  string
}

func main() {
	programArgs, err := parseProgramArguments()
	if err != nil {
		log.Fatalf("Error parsing program arguments: %v", err)
	}

	log.Printf("Reading program from file %v", programArgs.FilePath)
	programCode, err := os.ReadFile(programArgs.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	c := createCompiler()
	compiledCode, err := c.Compile(compiler.Args{
		SourceCode:            string(programCode),
		ShouldShowLogs:        true,
		ShouldShowProgramTree: false,
	})
	if err != nil {
		log.Fatalf("Error compiling program: %v", err)
	}

	err = os.WriteFile(programArgs.OutputPath, []byte(compiledCode), 0644)
	if err != nil {
		log.Fatalf("Error writing compiled code to file: %v", err)
	}
	log.Printf("Program compiled successfully")
}

func parseProgramArguments() (ProgramArguments, error) {
	args := os.Args

	if len(args) != 3 {
		return ProgramArguments{}, fmt.Errorf("expected 1 argument, got %v", len(args)-1)
	}

	programArgs := ProgramArguments{
		ProgramPath: args[0],
		FilePath:    args[1],
		OutputPath:  args[2],
	}

	return programArgs, nil
}

func createCompiler() compiler.Compiler {
	lexer := lexerModule.NewLexer()
	parser := newParser()
	emitter := emiterModule.NewCEmitter(
		emiterModule.NewCTokenEmitter(),
	)

	return compiler.NewCompiler(
		lexer,
		parser,
		emitter,
		log.Printf,
		utils.PrintProgramTree,
	)
}

func newParser() parserModule.Parser {
	var expressionParser parserModule.NodeParser

	factorParser := factor.NewFactorParser(expressionParser)
	termParser := term.NewTermParser(factorParser)
	expressionParser = expression.NewExpressionParser(termParser)
	expressionListParser := expressionList.NewExpressionListParser(expressionParser)
	relopParser := relop.NewRelopParser()
	varListParser := varList.NewVarListParser()
	statementParser := statement.NewStatementParser(expressionListParser, expressionParser, relopParser, varListParser)
	lineParser := line.NewLineParser(statementParser)

	return parserModule.NewParser(lineParser)
}
