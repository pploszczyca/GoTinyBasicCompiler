package expressionList

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"errors"
	"reflect"
	"testing"
)

func TestExpressionListParser_Parse(t *testing.T) {
	t.Run("returns error when index is out of range", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{})
		elp := NewExpressionListParser(&testutils.FakeNodeParser{})

		_, err := elp.Parse(&iterator)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		if err.Error() != "tokens index out of range" {
			t.Errorf("Expected tokens index out of range, got %v", err)
		}
	})

	t.Run("returns expression node with string token", func(t *testing.T) {
		stringToken := domain.Token{Type: domain.String}
		expectedExpressionListNode := domain.Node{
			Type: domain.ExpressionListNode,
			Children: []*domain.Node{
				{Token: stringToken},
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{stringToken, {Type: domain.Cr}})

		elp := NewExpressionListParser(&testutils.FakeNodeParser{
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
		fakeExpressionParser := testutils.FakeNodeParser{
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
		tokens := []domain.Token{identifierToken, {Type: domain.Cr}}
		fakeExpressionParser := testutils.FakeNodeParser{
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
			t.Errorf("Expected %v, got %v", expectedExpressionListNode, expressionListNode)
		}
	})

	t.Run("returns expression list node with string and expression node separated by comma", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		expressionNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionListNode := &domain.Node{
			Type: domain.ExpressionListNode,
			Children: []*domain.Node{
				{Token: domain.Token{Type: domain.String}},
				{Token: domain.Token{Type: domain.Comma}},
				expressionNode,
			},
		}
		tokens := []domain.Token{{Type: domain.String}, {Type: domain.Comma}, identifierToken, {Type: domain.Cr}}
		fakeExpressionParser := testutils.FakeNodeParser{
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
			t.Errorf("Expected %v, got %v", expectedExpressionListNode, expressionListNode)
		}
	})
}
