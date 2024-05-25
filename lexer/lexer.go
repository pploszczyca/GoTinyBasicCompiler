package lexer

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
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

	for i, line := range lines {
		tokens, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		if i != len(lines)-1 {
			tokens = append(tokens, domain.Token{Type: domain.Cr})
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
		log.Printf("token: %v", token)
		if err != nil && err.Error() != "empty token" {
			log.Printf("error: %v", err)
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func parseValue(value string) (domain.Token, error) {
	if value == "" {
		return domain.Token{}, fmt.Errorf("empty token")
	}
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
	if len(s) != 1 {
		return false
	}
	r := rune(s[0])
	return unicode.IsUpper(r) && unicode.IsLetter(r)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isString(s string) bool {
	return strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")
}
