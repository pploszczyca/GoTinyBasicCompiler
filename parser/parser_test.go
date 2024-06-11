package parser

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"fmt"
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	t.Run("returns program tree when tokens are valid", func(t *testing.T) {
		fakeLineParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{}, nil
			},
		}
		tokens := []domain.Token{{}, {}, {}}
		expectedNodes := []*domain.Node{{}, {}, {}}

		p := NewParser(fakeLineParser)
		programTree, err := p.Parse(tokens)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(programTree.Nodes, expectedNodes) {
			t.Errorf("Expected %v, got %v", expectedNodes, programTree.Nodes)
		}
	})

	t.Run("returns error when line parser returns error", func(t *testing.T) {
		parserError := fmt.Errorf("parse error")
		expectedError := fmt.Errorf("error parsing line 1: %v", parserError)
		fakeLineParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return nil, parserError
			},
		}
		tokens := []domain.Token{{}}

		p := NewParser(fakeLineParser)
		_, err := p.Parse(tokens)

		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error: %v, got %v", expectedError, err)
		}
	})

	t.Run("returns empty program tree when no tokens provided", func(t *testing.T) {
		var tokens []domain.Token

		p := NewParser(&testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{}, nil
			},
		})
		programTree, err := p.Parse(tokens)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(programTree.Nodes) != 0 {
			t.Errorf("Expected 0 nodes, got %d", len(programTree.Nodes))
		}
	})
}
