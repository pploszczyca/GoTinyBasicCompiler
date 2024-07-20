package internal

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/parser"
	"fmt"
)

type StepFunc func() error

func ParseSteps(steps []StepFunc) error {
	for _, step := range steps {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

func ExpectAndAddMatchingToken(
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
	expectedTokenType domain.TokenType,
) error {
	token, err := iterator.Current()
	if err != nil {
		return err
	}
	if token.Type != expectedTokenType {
		return fmt.Errorf("expected %s", expectedTokenType)
	}
	statementNode.AddChildToken(token)
	iterator.Next()

	return nil
}

func ParseAndAddNode(
	iterator *domain.TokenIterator,
	statementNode *domain.Node,
	parser parser.NodeParser,
) error {
	node, err := parser.Parse(iterator)
	if err != nil {
		return err
	}
	statementNode.AddChild(node)
	return nil
}
