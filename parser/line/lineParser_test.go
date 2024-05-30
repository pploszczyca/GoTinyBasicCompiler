package line

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"errors"
	"reflect"
	"testing"
)

func TestLineParser_Parse(t *testing.T) {
	t.Run("parses line with number and statement", func(t *testing.T) {
		identifierNode := &domain.Node{Type: domain.IdentifierNode}
		fakeStatementParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return identifierNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedLineNode := &domain.Node{
			Type: domain.LineNode,
			Children: []*domain.Node{
				{Type: domain.NumberNode, Token: tokens[0]},
				identifierNode,
			},
		}

		lp := NewLineParser(fakeStatementParser)
		lineNode, err := lp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(lineNode, expectedLineNode) {
			t.Errorf("Expected %v, got %v", expectedLineNode, lineNode)
		}
	})

	t.Run("returns error when statement parser returns error", func(t *testing.T) {
		expectedError := errors.New("parse error")
		fakeStatementParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		lp := NewLineParser(fakeStatementParser)
		_, err := lp.Parse(&iterator)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when index is out of range", func(t *testing.T) {
		fakeStatementParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{}, nil
			},
		}
		tokens := []domain.Token{{Type: domain.Number}}
		iterator := domain.NewTokenIterator(tokens)

		lp := NewLineParser(fakeStatementParser)
		_, err := lp.Parse(&iterator)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns error when no CR token at the end", func(t *testing.T) {
		fakeStatementParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		lp := NewLineParser(fakeStatementParser)
		_, err := lp.Parse(&iterator)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
