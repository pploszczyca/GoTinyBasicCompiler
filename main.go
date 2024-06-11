package GoTinyBasicCompiler

import (
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

	compiledCode, err := compile(string(programCode))
	if err != nil {
		log.Fatalf("Error compiling program: %v", err)
	}

	err = os.WriteFile(programArgs.OutputPath, []byte(compiledCode), 0644)
	if err != nil {
		log.Fatalf("Error writing compiled code to file: %v", err)
	}
	log.Printf("Program compiled successfully")
}

func compile(programCode string) (string, error) {
	lexer := lexerModule.NewLexer()
	parser := newParser()

	tokens, err := lexer.Lex(programCode)
	if err != nil {
		return "", fmt.Errorf("error lexing program: %v", err)
	}

	programTree, err := parser.Parse(tokens)
	if err != nil {
		return "", fmt.Errorf("error parsing program: %v", err)
	}

	fmt.Printf("Program tree:\n")
	utils.PrintProgramTree(&programTree)

	return "", nil
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
