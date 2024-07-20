package lexer

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
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
	return l.sequentialLex(inputCode)
}

func (l *lexer) sequentialLex(inputCode string) ([]domain.Token, error) {
	lines := strings.Split(inputCode, "\n")
	var result []domain.Token

	for i, line := range lines {
		tokens, err := lexLine(line)
		if err != nil {
			return nil, fmt.Errorf("error lexing line %d: %v", i+1, err)
		}
		if i != len(lines)-1 {
			tokens = append(tokens, domain.Token{Type: domain.Cr})
		}

		result = append(result, tokens...)
	}

	result = append(result, domain.Token{Type: domain.Eof})

	return result, nil
}

func lexLine(line string) ([]domain.Token, error) {
	var tokens []domain.Token
	currentIndex := 0
	for currentIndex < len(line) {
		char := line[currentIndex]

		if char == '"' {
			token, newIndex, err := readStringToken(line, currentIndex)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			currentIndex = newIndex
		} else if unicode.IsDigit(rune(char)) {
			token, newIndex, err := readNumberToken(line, currentIndex)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			currentIndex = newIndex
		} else if char != ' ' {
			token, newIndex, err := readAnotherToken(line, currentIndex)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			currentIndex = newIndex
		}

		if char == '\n' {
			tokens = append(tokens, domain.Token{Type: domain.Cr})
		}

		currentIndex++
	}

	return tokens, nil
}

func readStringToken(line string, currentIndex int) (domain.Token, int, error) {
	result := "\""
	currentIndex++
	for currentIndex < len(line) {
		result += string(line[currentIndex])
		if line[currentIndex] == '"' {
			return domain.Token{Type: domain.String, Value: result}, currentIndex, nil
		}
		currentIndex++
	}
	return domain.Token{}, currentIndex, fmt.Errorf("unterminated string")
}

func readNumberToken(line string, currentIndex int) (domain.Token, int, error) {
	result := string(line[currentIndex])
	currentIndex++
	for currentIndex < len(line) {
		char := line[currentIndex]

		if unicode.IsDigit(rune(char)) {
			result += string(char)
			currentIndex++
		} else if char == ' ' || char == '\n' {
			return domain.Token{Type: domain.Number, Value: result}, currentIndex, nil
		} else {
			return domain.Token{}, currentIndex, fmt.Errorf("invalid number")
		}
	}

	return domain.Token{Type: domain.Number, Value: result}, currentIndex, nil
}

func readAnotherToken(line string, currentIndex int) (domain.Token, int, error) {
	result := string(line[currentIndex])
	currentIndex++

	for currentIndex < len(line) {
		char := line[currentIndex]
		if char == ' ' || char == '\n' || char == ',' {
			if char == ',' {
				currentIndex--
			}

			return lexAnother(result, currentIndex)
		}
		result += string(char)
		currentIndex++
	}

	return lexAnother(result, currentIndex)
}

func lexAnother(result string, currentIndex int) (domain.Token, int, error) {
	keywordToken, err := parseToKeyword(result)
	if err != nil {
		operatorToken, err := parseToOperator(result)
		if err != nil {
			delimeterOperator, err := parseToDelimiter(result)
			if err != nil {
				identifierToken, err := parseIdentifier(result)
				if err != nil {
					return domain.Token{}, currentIndex, err
				}
				return identifierToken, currentIndex, nil
			}
			return delimeterOperator, currentIndex, nil
		}
		return operatorToken, currentIndex, nil
	}
	return keywordToken, currentIndex, nil
}

func parseToKeyword(value string) (domain.Token, error) {
	tokenType, ok := map[string]domain.TokenType{
		"PRINT": domain.Print, "IF": domain.If, "THEN": domain.Then, "GOTO": domain.Goto,
		"INPUT": domain.Input, "LET": domain.Let, "GOSUB": domain.Gosub, "RETURN": domain.Return,
		"CLEAR": domain.Clear, "LIST": domain.List, "RUN": domain.Run, "END": domain.End,
		"WHILE": domain.While, "WEND": domain.Wend, "FOR": domain.For, "TO": domain.To, "NEXT": domain.Next,
	}[value]

	if !ok {
		return domain.Token{}, fmt.Errorf("invalid keyword")
	}
	return domain.Token{Type: tokenType}, nil
}

func parseToOperator(value string) (domain.Token, error) {
	switch value {
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
	}

	return domain.Token{}, fmt.Errorf("invalid operator")
}

func parseToDelimiter(value string) (domain.Token, error) {
	switch value {
	case ",":
		return domain.Token{Type: domain.Comma}, nil
	case ";":
		return domain.Token{Type: domain.Semicolon}, nil
	case "(":
		return domain.Token{Type: domain.LParen}, nil
	case ")":
		return domain.Token{Type: domain.RParen}, nil
	}

	return domain.Token{}, fmt.Errorf("invalid delimiter")
}

func parseIdentifier(value string) (domain.Token, error) {
	if len(value) != 1 {
		return domain.Token{}, fmt.Errorf("invalid token: %s", value)
	}

	return domain.Token{Type: domain.Identifier, Value: value}, nil
}
