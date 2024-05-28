package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type statementParser struct {
	expressionListParser parser.NodeParser
}

func NewStatementParser(
	expressionListParser parser.NodeParser,
) parser.NodeParser {
	return &statementParser{
		expressionListParser: expressionListParser,
	}
}

func (s statementParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	statementNode := domain.Node{Type: domain.StatementNode}

	switch tokens[currentIndex].Type {
	case domain.Print:
		statementNode.AddChild(&domain.Node{Token: tokens[currentIndex]})
		currentIndex++
		expressionListNode, newIndex, err := s.expressionListParser.Parse(tokens, currentIndex)
		if err != nil {
			return nil, newIndex, err
		}
		currentIndex = newIndex
		statementNode.AddChild(expressionListNode)

	// TODO: Implement parsing of other statements

	default:
		return nil, currentIndex, fmt.Errorf("unexpected statement: %v", tokens[currentIndex].Type)
	}

	return &statementNode, currentIndex, nil
}
