package line

import (
	"GoTinyBasicCompiler/domain"
	"errors"
	"reflect"
	"testing"
)

func TestLineParser_Parse(t *testing.T) {
	t.Run("parses line with number and statement", func(t *testing.T) {
		identifierNode := &domain.Node{Type: domain.IdentifierNode}
		fakeStatementParser := &fakeStatementParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return identifierNode, currentIndex + 1, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
			{Type: domain.Cr},
		}
		expectedLineNode := &domain.Node{
			Type: domain.LineNode,
			Children: []*domain.Node{
				{Type: domain.NumberNode, Token: tokens[0]},
				identifierNode,
			},
		}

		lp := NewLineParser(fakeStatementParser)
		lineNode, index, err := lp.Parse(tokens, 0)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(lineNode, expectedLineNode) {
			t.Errorf("Expected %v, got %v", expectedLineNode, lineNode)
		}
		if index != 3 {
			t.Errorf("Expected index 3, got %d", index)
		}
	})

	t.Run("returns error when statement parser returns error", func(t *testing.T) {
		expectedError := errors.New("parse error")
		fakeStatementParser := &fakeStatementParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return nil, 0, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
			{Type: domain.Cr},
		}

		lp := NewLineParser(fakeStatementParser)
		_, _, err := lp.Parse(tokens, 0)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when index is out of range", func(t *testing.T) {
		fakeStatementParser := &fakeStatementParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return &domain.Node{}, currentIndex + 1, nil
			},
		}
		tokens := []domain.Token{{Type: domain.Number}}

		lp := NewLineParser(fakeStatementParser)
		_, _, err := lp.Parse(tokens, 0)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns error when no CR token at the end", func(t *testing.T) {
		fakeStatementParser := &fakeStatementParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return &domain.Node{}, currentIndex + 1, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Number},
			{Type: domain.Identifier},
		}

		lp := NewLineParser(fakeStatementParser)
		_, _, err := lp.Parse(tokens, 0)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

type fakeStatementParser struct {
	ParseMock func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error)
}

func (f fakeStatementParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	return f.ParseMock(tokens, currentIndex)
}
