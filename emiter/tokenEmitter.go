package emiter

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
)

type TokenEmitter interface {
	Emit(token domain.Token) (string, error)
}

type cTokenEmitter struct {
}

func NewCTokenEmitter() TokenEmitter {
	return &cTokenEmitter{}
}

func (c *cTokenEmitter) Emit(token domain.Token) (string, error) {
	switch token.Type {
	case domain.Number:
		return token.Value, nil
	case domain.Identifier:
		return token.Value, nil
	case domain.String:
		return token.Value, nil
	case domain.Print:
		return "printf", nil
	case domain.Input:
		return "scanf", nil
	case domain.Plus:
		return "+", nil
	case domain.Minus:
		return "-", nil
	case domain.Multiply:
		return "*", nil
	case domain.Divide:
		return "/", nil
	case domain.Equal:
		return "=", nil
	case domain.LessThan:
		return "<", nil
	case domain.GreaterThan:
		return ">", nil
	case domain.LessThanOrEqual:
		return "<=", nil
	case domain.GreaterThanOrEqual:
		return ">=", nil
	case domain.NotEqual:
		return "!=", nil
	case domain.Comma:
		return ", ", nil
	case domain.Semicolon:
		return ";", nil
	case domain.LParen:
		return "(", nil
	case domain.RParen:
		return ")", nil
	case domain.End:
		return "return 0", nil
	case domain.Let:
		return "int", nil
	case domain.If:
		return "if", nil
	case domain.Goto:
		return "goto", nil
	case domain.While:
		return "while", nil
	case domain.Wend:
		return "}", nil
	case domain.For:
		return "for", nil
	case domain.Next:
		return "}", nil
	case domain.To:
		return "", nil
	case domain.Gosub:
		return "goto", nil
	}

	return "", fmt.Errorf("Unknown token type: %s\n", token.Type)
}
