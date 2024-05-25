package lexer

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"strconv"
	"strings"
)

type Lexer interface {
	Lex(inputCode string) ([]domain.Token, error)
}

type lexer struct {
}

func NewLexer() Lexer {
	return &lexer{}
}

func (l *lexer) Lex(inputCode string) ([]domain.Token, error) {
	lines := strings.Split(inputCode, "\n")
	var result []domain.Token

	for _, line := range lines {
		tokens, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		result = append(result, tokens...)
	}

	result = append(result, domain.Token{Type: domain.Eof})

	return result, nil
}

func parseLine(line string) ([]domain.Token, error) {
	var tokens []domain.Token
	values := strings.Split(line, "\t")

	for _, value := range values {
		token, err := parseValue(value)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	tokens = append(tokens, domain.Token{Type: domain.Cr})

	return tokens, nil
}

func parseValue(value string) (domain.Token, error) {
	if isNumber(value) {
		return domain.Token{Type: domain.Number, Value: value}, nil
	}
	if isString(value) {
		return domain.Token{Type: domain.String, Value: value}, nil
	}

	switch value {
	case "PRINT":
		return domain.Token{Type: domain.Print}, nil
	case "IF":
		return domain.Token{Type: domain.If}, nil
	case "THEN":
		return domain.Token{Type: domain.Then}, nil
	case "GOTO":
		return domain.Token{Type: domain.Goto}, nil
	case "INPUT":
		return domain.Token{Type: domain.Input}, nil
	case "LET":
		return domain.Token{Type: domain.Let}, nil
	case "GOSUB":
		return domain.Token{Type: domain.Gosub}, nil
	case "RETURN":
		return domain.Token{Type: domain.Return}, nil
	case "CLEAR":
		return domain.Token{Type: domain.Clear}, nil
	case "LIST":
		return domain.Token{Type: domain.List}, nil
	case "RUN":
		return domain.Token{Type: domain.Run}, nil
	case "END":
		return domain.Token{Type: domain.End}, nil
	case "+":
		return domain.Token{Type: domain.Plus}, nil
	case "-":
		return domain.Token{Type: domain.Minus}, nil
	case "*":
		return domain.Token{Type: domain.Multiply}, nil
	case "/":
		return domain.Token{Type: domain.Divide}, nil
	case "=":
		return domain.Token{Type: domain.Equal}, nil
	case "<":
		return domain.Token{Type: domain.LessThan}, nil
	case ">":
		return domain.Token{Type: domain.GreaterThan}, nil
	case "<=":
		return domain.Token{Type: domain.LessThanOrEqual}, nil
	case ">=":
		return domain.Token{Type: domain.GreaterThanOrEqual}, nil
	case "<>":
		return domain.Token{Type: domain.NotEqual}, nil
	case ",":
		return domain.Token{Type: domain.Comma}, nil
	case ";":
		return domain.Token{Type: domain.Semicolon}, nil
	case "(":
		return domain.Token{Type: domain.LParen}, nil
	case ")":
		return domain.Token{Type: domain.RParen}, nil
	}

	if isIdentifier(value) {
		return domain.Token{Type: domain.Identifier, Value: value}, nil
	}

	return domain.Token{}, fmt.Errorf("invalid token: %s", value)
}

func isIdentifier(s string) bool {
	return !isNumber(s) && !isString(s)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isString(s string) bool {
	return strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")
}
