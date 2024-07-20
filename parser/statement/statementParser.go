package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"GoTinyBasicCompiler/parser/internal"
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
	case domain.For:
		parseErr = s.parseFor(token, iterator, &statementNode)
	case domain.Next:
		parseErr = s.parseNext(token, iterator, &statementNode)
	case domain.Return, domain.Clear, domain.List, domain.Run, domain.End, domain.Wend, domain.To:
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

	return internal.ParseAndAddNode(iterator, statementNode, s.expressionListParser)
}

func (s statementParser) parseIf(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	if err := internal.ParseSteps([]internal.StepFunc{
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.relopParser) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Then) },
	}); err != nil {
		return err
	}

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

	return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser)
}

func (s statementParser) parseInput(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ParseAndAddNode(iterator, statementNode, s.varListParser)
}

func (s statementParser) parseLet(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ParseSteps([]internal.StepFunc{
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Identifier) },
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Equal) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
	})
}

func (s statementParser) parseGosub(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser)
}

func (s statementParser) parseWhile(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ParseSteps([]internal.StepFunc{
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.relopParser) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
	})
}

func (s statementParser) parseFor(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ParseSteps([]internal.StepFunc{
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Identifier) },
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Equal) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
		func() error { return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.To) },
		func() error { return internal.ParseAndAddNode(iterator, statementNode, s.expressionParser) },
	})
}

func (s statementParser) parseNext(
	token domain.Token,
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
) error {
	statementNode.AddChildToken(token)
	iterator.Next()

	return internal.ExpectAndAddMatchingToken(iterator, statementNode, domain.Identifier)
}
