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
	keywords   map[string]domain.TokenType
	operators  map[string]domain.TokenType
	delimiters map[string]domain.TokenType
}

func NewLexer() Lexer {
	keywords := map[string]domain.TokenType{
		"PRINT":  domain.Print,
		"IF":     domain.If,
		"THEN":   domain.Then,
		"GOTO":   domain.Goto,
		"INPUT":  domain.Input,
		"LET":    domain.Let,
		"GOSUB":  domain.Gosub,
		"RETURN": domain.Return,
		"CLEAR":  domain.Clear,
		"LIST":   domain.List,
		"RUN":    domain.Run,
		"END":    domain.End,
		"WHILE":  domain.While,
		"WEND":   domain.Wend,
		"FOR":    domain.For,
		"TO":     domain.To,
		"NEXT":   domain.Next,
	}

	operators := map[string]domain.TokenType{
		"+":  domain.Plus,
		"-":  domain.Minus,
		"*":  domain.Multiply,
		"/":  domain.Divide,
		"=":  domain.Equal,
		"<":  domain.LessThan,
		">":  domain.GreaterThan,
		"<=": domain.LessThanOrEqual,
		">=": domain.GreaterThanOrEqual,
		"<>": domain.NotEqual,
	}

	delimiters := map[string]domain.TokenType{
		",": domain.Comma,
		";": domain.Semicolon,
		"(": domain.LParen,
		")": domain.RParen,
	}

	return &lexer{
		keywords:   keywords,
		operators:  operators,
		delimiters: delimiters,
	}
}

func (l *lexer) Lex(inputCode string) ([]domain.Token, error) {
	return l.sequentialLex(inputCode)
}

func (l *lexer) sequentialLex(inputCode string) ([]domain.Token, error) {
	lines := strings.Split(inputCode, "\n")
	var result []domain.Token

	for i, line := range lines {
		tokens, err := l.lexLine(line)
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

func (l *lexer) lexLine(line string) ([]domain.Token, error) {
	var tokens []domain.Token
	currentIndex := 0
	for currentIndex < len(line) {
		char := line[currentIndex]

		if char == '"' {
			token, newIndex, err := l.readStringToken(line, currentIndex)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			currentIndex = newIndex
		} else if unicode.IsDigit(rune(char)) {
			token, newIndex, err := l.readNumberToken(line, currentIndex)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			currentIndex = newIndex
		} else if char != ' ' {
			token, newIndex, err := l.readAnotherToken(line, currentIndex)
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

func (l *lexer) readStringToken(line string, currentIndex int) (domain.Token, int, error) {
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

func (l *lexer) readNumberToken(line string, currentIndex int) (domain.Token, int, error) {
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

func (l *lexer) readAnotherToken(line string, currentIndex int) (domain.Token, int, error) {
	result := string(line[currentIndex])
	currentIndex++

	for currentIndex < len(line) {
		char := line[currentIndex]
		if char == ' ' || char == '\n' || char == ',' {
			if char == ',' {
				currentIndex--
			}

			return l.lexAnother(result, currentIndex)
		}
		result += string(char)
		currentIndex++
	}

	return l.lexAnother(result, currentIndex)
}

func (l *lexer) lexAnother(result string, currentIndex int) (domain.Token, int, error) {
	if token, err := l.parseToKeyword(result); err == nil {
		return token, currentIndex, nil
	}
	if token, err := l.parseToOperator(result); err == nil {
		return token, currentIndex, nil
	}
	if token, err := l.parseToDelimiter(result); err == nil {
		return token, currentIndex, nil
	}

	token, err := l.parseIdentifier(result)

	return token, currentIndex, err
}

func (l *lexer) parseToKeyword(value string) (domain.Token, error) {
	if tokenType, ok := l.keywords[value]; ok {
		return domain.Token{Type: tokenType}, nil
	}

	return domain.Token{}, fmt.Errorf("invalid keyword")
}

func (l *lexer) parseToOperator(value string) (domain.Token, error) {
	if tokenType, ok := l.operators[value]; ok {
		return domain.Token{Type: tokenType}, nil
	}
	return domain.Token{}, fmt.Errorf("invalid operator")
}

func (l *lexer) parseToDelimiter(value string) (domain.Token, error) {
	if tokenType, ok := l.delimiters[value]; ok {
		return domain.Token{Type: tokenType}, nil
	}
	return domain.Token{}, fmt.Errorf("invalid delimiter")
}

func (l *lexer) parseIdentifier(value string) (domain.Token, error) {
	if len(value) != 1 {
		return domain.Token{}, fmt.Errorf("invalid token: %s", value)
	}

	return domain.Token{Type: domain.Identifier, Value: value}, nil
}
