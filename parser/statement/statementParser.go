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

func (s statementParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	statementNode := domain.Node{Type: domain.StatementNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	switch token.Type {
	case domain.Print:
		statementNode.AddChildToken(token)
		iterator.Next()
		expressionListNode, err := s.expressionListParser.Parse(iterator)
		if err != nil {
			return nil, err
		}
		statementNode.AddChild(expressionListNode)

	// TODO: Implement parsing of other statements

	default:
		return nil, fmt.Errorf("unexpected statement: %v", token.Type)
	}

	return &statementNode, nil
}
