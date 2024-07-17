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
	varListParser        parser.NodeParser
}

func NewStatementParser(
	expressionListParser parser.NodeParser,
	expressionParser parser.NodeParser,
	relopParser parser.NodeParser,
	varListParser parser.NodeParser,
) parser.NodeParser {
	return &statementParser{
		expressionListParser: expressionListParser,
		expressionParser:     expressionParser,
		relopParser:          relopParser,
		varListParser:        varListParser,
	}
}

func (s statementParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	statementNode := domain.Node{Type: domain.StatementNode}

	token, err := iterator.Current()
	if err != nil {
		return nil, err
	}

	var parseErr error
	switch token.Type {
	case domain.Print:
		parseErr = s.parsePrint(token, iterator, &statementNode)
	case domain.If:
		parseErr = s.parseIf(token, iterator, &statementNode)
	case domain.Goto:
		parseErr = s.parseGoto(token, iterator, &statementNode)
	case domain.Input:
		parseErr = s.parseInput(token, iterator, &statementNode)
	case domain.Let:
		parseErr = s.parseLet(token, iterator, &statementNode)
	case domain.Gosub:
		parseErr = s.parseGosub(token, iterator, &statementNode)
	case domain.While:
		parseErr = s.parseWhile(token, iterator, &statementNode)
	case domain.Return, domain.Clear, domain.List, domain.Run, domain.End, domain.Wend:
		statementNode.AddChildToken(token)
		iterator.Next()
	default:
		return nil, fmt.Errorf("unexpected statement: %v", token.Type)
	}

	if parseErr != nil {
		return nil, parseErr
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

func (s statementParser) parseInput(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	varListNode, err := s.varListParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(varListNode)

	return nil
}

func (s statementParser) parseLet(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	token, err := iterator.Current()
	if err != nil {
		return err
	}
	if token.Type != domain.Identifier {
		return fmt.Errorf("expected identifier")
	}
	statementNode.AddChildToken(token)
	iterator.Next()

	token, err = iterator.Current()
	if err != nil {
		return err
	}
	if token.Type != domain.Equal {
		return fmt.Errorf("expected equal")
	}
	statementNode.AddChildToken(token)
	iterator.Next()

	expressionNode, err := s.expressionParser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(expressionNode)

	return nil
}

func (s statementParser) parseGosub(
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

func (s statementParser) parseWhile(
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

	return nil
}
