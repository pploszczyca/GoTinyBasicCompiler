package internal

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestParseSteps(t *testing.T) {
	t.Run("returns error when one step fails", func(t *testing.T) {
		expectedError := fmt.Errorf("error")
		steps := []StepFunc{
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
			func() error { return expectedError },
			func() error { return nil },
		}

		err := ParseSteps(steps)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns nil when all steps succeed", func(t *testing.T) {
		steps := []StepFunc{
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
			func() error { return nil },
		}

		err := ParseSteps(steps)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}

func TestExpectAndAddMatchingToken(t *testing.T) {
	t.Run("returns error when current token is not present", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{})
		expectedError := "tokens index out of range"

		err := ExpectAndAddMatchingToken(&iterator, &domain.Node{}, domain.Identifier)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when current token is not the expected token", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Number}})
		expectedError := "expected Identifier"

		err := ExpectAndAddMatchingToken(&iterator, &domain.Node{}, domain.Identifier)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("adds token to node and moves iterator to next token", func(t *testing.T) {
		firstToken := domain.Token{Type: domain.Number}
		secondToken := domain.Token{Type: domain.Identifier}
		iterator := domain.NewTokenIterator([]domain.Token{firstToken, secondToken})
		expectedNode := domain.Node{Token: firstToken}

		err := ExpectAndAddMatchingToken(&iterator, &domain.Node{}, domain.Number)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		token, err := iterator.Current()
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if token != secondToken {
			t.Errorf("Expected %v, got %v", expectedNode, token)
		}
	})
}

func TestParseAndAddNode(t *testing.T) {
	t.Run("returns error when parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("error")
		fakeParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{})

		err := ParseAndAddNode(&iterator, &domain.Node{}, fakeParser)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("adds node to parent node", func(t *testing.T) {
		newNode := domain.Node{Type: domain.ExpressionNode}
		fakeParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return &newNode, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{})
		statementNode := &domain.Node{Type: domain.ExpressionNode}
		expectedNode := &domain.Node{Type: domain.ExpressionNode, Children: []*domain.Node{&newNode}}

		err := ParseAndAddNode(&iterator, statementNode, fakeParser)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(expectedNode, statementNode) {
			t.Errorf("Expected %v, got %v", expectedNode, &domain.Node{})
		}
	})
}
