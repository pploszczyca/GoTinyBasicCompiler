package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/expressionList"
	"fmt"
)

type statementParser struct {
	expressionListParser parser.NodeParser
}

func NewStatementParser() parser.NodeParser {
	return &statementParser{
		expressionListParser: expressionList.NewExpressionListParser(),
	}
}

func (s statementParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	statementNode := domain.Node{Type: domain.StatementNode}

	switch tokens[currentIndex].Type {
	case domain.Print:
		statementNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
		expressionListNode, currentIndex, err := s.expressionListParser.Parse(tokens, currentIndex)
		if err != nil {
			return nil, currentIndex, err
		}
		statementNode.AddChild(expressionListNode)

	default:
		return nil, currentIndex, fmt.Errorf("unexpected statement: %v", tokens[currentIndex].Type)
	}

	return nil, -1, nil
}
