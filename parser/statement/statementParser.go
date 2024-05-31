package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type statementParser struct {
	expressionListParser parser.NodeParser
	expressionParser     parser.NodeParser
	relopParser          parser.NodeParser
}

func NewStatementParser(
	expressionListParser parser.NodeParser,
	expressionParser parser.NodeParser,
	relopParser parser.NodeParser,
) parser.NodeParser {
	return &statementParser{
		expressionListParser: expressionListParser,
		expressionParser:     expressionParser,
		relopParser:          relopParser,
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
		err := s.parsePrint(token, iterator, &statementNode)
		if err != nil {
			return nil, err
		}
	case domain.If:
		err := s.parseIf(token, iterator, &statementNode)
		if err != nil {
			return nil, err
		}

	case domain.Goto:
		err := s.parseGoto(token, iterator, &statementNode)
		if err != nil {
			return nil, err
		}

	// TODO: Implement parsing of other statements

	default:
		return nil, fmt.Errorf("unexpected statement: %v", token.Type)
	}

	return &statementNode, nil
}

func (s statementParser) parsePrint(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()
	expressionListNode, err := s.expressionListParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(expressionListNode)
	return nil
}

func (s statementParser) parseIf(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	expressionNode, err := s.expressionParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(expressionNode)

	relopNode, err := s.relopParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(relopNode)

	expressionNode, err = s.expressionParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(expressionNode)

	token, err = iterator.Current()
	if err != nil {
		return err
	}
	if token.Type != domain.Then {
		return fmt.Errorf("expected THEN")
	}
	statementNode.AddChildToken(token)
	iterator.Next()

	ifStatementNode, err := s.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(ifStatementNode)

	return nil
}

func (s statementParser) parseGoto(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	expressionNode, err := s.expressionParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(expressionNode)

	return nil
}
