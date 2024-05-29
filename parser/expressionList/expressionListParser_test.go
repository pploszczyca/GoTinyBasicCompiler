package expressionList

import (
	"GoTinyBasicCompiler/domain"
	"errors"
	"reflect"
	"testing"
)

func TestExpressionListParser_Parse(t *testing.T) {
	t.Run("returns expression node with string token", func(t *testing.T) {
		stringToken := domain.Token{Type: domain.String}
		expectedExpressionListNode := domain.Node{
			Type: domain.ExpressionListNode,
			Children: []*domain.Node{
				{Token: stringToken},
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{stringToken})

		elp := NewExpressionListParser(&fakeExpressionParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{}, nil
			},
		})
		expressionListNode, err := elp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionListNode.Type != expressionListNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionListNode.Type, expressionListNode.Type)
		}
		if !reflect.DeepEqual(*expressionListNode.Children[0], *expectedExpressionListNode.Children[0]) {
			t.Errorf("Expected %v, got %v", *expectedExpressionListNode.Children[0], *expressionListNode.Children[0])
		}
	})

	t.Run("returns error when expression parser returns error", func(t *testing.T) {
		expectedError := errors.New("parse error")
		fakeExpressionParser := fakeExpressionParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier}})

		elp := NewExpressionListParser(&fakeExpressionParser)
		_, err := elp.Parse(&iterator)

		if !errors.Is(err, expectedError) {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns expression list node with expression node", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		expressionNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionListNode := &domain.Node{
			Type: domain.ExpressionListNode,
			Children: []*domain.Node{
				expressionNode,
			},
		}
		tokens := []domain.Token{identifierToken}
		fakeExpressionParser := fakeExpressionParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		iterator := domain.NewTokenIterator(tokens)

		elp := NewExpressionListParser(&fakeExpressionParser)
		expressionListNode, err := elp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(expressionListNode, expectedExpressionListNode) {
			println(expressionListNode.Children[0].Token.Type)
			println(expectedExpressionListNode.Children[0].Token.Type)
			t.Errorf("Expected %v, got %v", expectedExpressionListNode, expressionListNode)
		}
	})
}

type fakeExpressionParser struct {
	ParseMock func(iterator *domain.TokenIterator) (*domain.Node, error)
}

func (f *fakeExpressionParser) Parse(iterator *domain.TokenIterator) (*domain.Node, error) {
	return f.ParseMock(iterator)
}
