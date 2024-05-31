package expression

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"reflect"
	"testing"
)

func TestExpressionParser_Parse(t *testing.T) {
	t.Run("returns error when index is out of range", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{})
		ep := NewExpressionParser(&testutils.FakeNodeParser{})
		expectedError := "tokens index out of range"

		_, err := ep.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns expression node with term node", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		termNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionNode := &domain.Node{
			Type: domain.ExpressionNode,
			Children: []*domain.Node{
				termNode,
			},
		}
		tokens := []domain.Token{identifierToken, {Type: domain.Cr}}
		fakeTermParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return termNode, nil
			},
		}
		iterator := domain.NewTokenIterator(tokens)

		ep := NewExpressionParser(&fakeTermParser)
		expressionNode, err := ep.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionNode.Type != expressionNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Type, expressionNode.Type)
		}
		if !reflect.DeepEqual(expectedExpressionNode.Children[0], expressionNode.Children[0]) {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Children[0], expressionNode.Children[0])
		}
	})

	t.Run("returns expression node with term node and plus token", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		termNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionNode := &domain.Node{
			Type: domain.ExpressionNode,
			Children: []*domain.Node{
				{Token: domain.Token{Type: domain.Plus}},
				termNode,
			},
		}
		tokens := []domain.Token{{Type: domain.Plus}, identifierToken, {Type: domain.Cr}}
		fakeTermParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return termNode, nil
			},
		}
		iterator := domain.NewTokenIterator(tokens)

		ep := NewExpressionParser(&fakeTermParser)
		expressionNode, err := ep.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionNode.Type != expressionNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Type, expressionNode.Type)
		}
		if !reflect.DeepEqual(expectedExpressionNode.Children, expressionNode.Children) {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Children, expressionNode.Children)
		}
	})

	t.Run("returns expression node with term node and minus token", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		termNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionNode := &domain.Node{
			Type: domain.ExpressionNode,
			Children: []*domain.Node{
				{Token: domain.Token{Type: domain.Minus}},
				termNode,
			},
		}
		tokens := []domain.Token{{Type: domain.Minus}, identifierToken, {Type: domain.Cr}}
		fakeTermParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return termNode, nil
			},
		}
		iterator := domain.NewTokenIterator(tokens)

		ep := NewExpressionParser(&fakeTermParser)
		expressionNode, err := ep.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionNode.Type != expressionNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Type, expressionNode.Type)
		}
		if !reflect.DeepEqual(expectedExpressionNode.Children, expressionNode.Children) {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Children, expressionNode.Children)
		}
	})

	t.Run("returns expression node with multiple identifies", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		termNode := &domain.Node{Token: domain.Token{Type: domain.Identifier}}
		expectedExpressionNode := &domain.Node{
			Type: domain.ExpressionNode,
			Children: []*domain.Node{
				termNode,
				{Token: domain.Token{Type: domain.Plus}},
				termNode,
				{Token: domain.Token{Type: domain.Minus}},
				termNode,
			},
		}
		tokens := []domain.Token{identifierToken, {Type: domain.Plus}, identifierToken, {Type: domain.Minus}, identifierToken, {Type: domain.Cr}}
		fakeTermParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return termNode, nil
			},
		}
		iterator := domain.NewTokenIterator(tokens)

		ep := NewExpressionParser(&fakeTermParser)
		expressionNode, err := ep.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionNode.Type != expressionNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Type, expressionNode.Type)
		}
		if !reflect.DeepEqual(expectedExpressionNode.Children, expressionNode.Children) {
			t.Errorf("Expected %v, got %v", expectedExpressionNode.Children, expressionNode.Children)
		}
	})
}
