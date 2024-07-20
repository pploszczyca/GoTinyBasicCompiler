package factor

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"reflect"
	"testing"
)

func TestFactorParser_Parse(t *testing.T) {
	t.Run("returns factor node when factor is an identifier", func(t *testing.T) {
		identifierToken := domain.Token{Type: domain.Identifier}
		iterator := domain.NewTokenIterator([]domain.Token{identifierToken})
		expectedFactorNode := &domain.Node{Type: domain.FactorNode, Children: []*domain.Node{{Token: identifierToken}}}

		fp := NewFactorParser(&testutils.FakeNodeParser{})
		factorNode, err := fp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(factorNode, expectedFactorNode) {
			t.Errorf("Expected %v, got %v", expectedFactorNode, factorNode)
		}
	})

	t.Run("returns factor node when factor is a number", func(t *testing.T) {
		numberToken := domain.Token{Type: domain.Number}
		iterator := domain.NewTokenIterator([]domain.Token{numberToken})
		expectedFactorNode := &domain.Node{Type: domain.FactorNode, Children: []*domain.Node{{Token: numberToken}}}

		fp := NewFactorParser(&testutils.FakeNodeParser{})
		factorNode, err := fp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(factorNode, expectedFactorNode) {
			t.Errorf("Expected %v, got %v", expectedFactorNode, factorNode)
		}
	})

	t.Run("returns factor node when factor is an expression", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.LParen}, {Type: domain.Identifier}, {Type: domain.RParen}})
		expectedFactorNode := &domain.Node{Type: domain.FactorNode, Children: []*domain.Node{
			{Token: domain.Token{Type: domain.LParen}},
			expressionNode,
			{Token: domain.Token{Type: domain.RParen}},
		}}

		fp := NewFactorParser(fakeExpressionParser)
		factorNode, err := fp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(factorNode, expectedFactorNode) {
			t.Errorf("Expected %v, got %v", expectedFactorNode, factorNode)
		}
	})

	t.Run("returns error when index is out of range", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.LParen}, {Type: domain.Identifier}})
		expectedError := "tokens index out of range"

		fp := NewFactorParser(fakeExpressionParser)
		_, err := fp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when there is no right parenthesis", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.LParen}, {Type: domain.Identifier}, {Type: domain.Number}})
		expectedError := "expected RParen"

		fp := NewFactorParser(fakeExpressionParser)
		_, err := fp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when unexpected token", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Plus}})
		expectedError := "expected number, identifier or left parenthesis"

		fp := NewFactorParser(&testutils.FakeNodeParser{})
		_, err := fp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})
}
