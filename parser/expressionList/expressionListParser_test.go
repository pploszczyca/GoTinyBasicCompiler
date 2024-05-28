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

		elp := NewExpressionListParser(fakeExpressionParser{})
		expressionListNode, newIndex, err := elp.Parse([]domain.Token{stringToken}, 0)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if expectedExpressionListNode.Type != expressionListNode.Type {
			t.Errorf("Expected %v, got %v", expectedExpressionListNode.Type, expressionListNode.Type)
		}
		if !reflect.DeepEqual(*expressionListNode.Children[0], *expectedExpressionListNode.Children[0]) {
			t.Errorf("Expected %v, got %v", *expectedExpressionListNode.Children[0], *expressionListNode.Children[0])
		}
		if newIndex != 1 {
			t.Errorf("Expected index 1, got %d", newIndex)
		}
	})

	t.Run("returns error when expression parser returns error", func(t *testing.T) {
		expectedError := errors.New("parse error")
		fakeExpressionParser := fakeExpressionParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return nil, 0, expectedError
			},
		}

		elp := NewExpressionListParser(fakeExpressionParser)
		_, _, err := elp.Parse([]domain.Token{{Type: domain.Identifier}}, 0)

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
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return expressionNode, currentIndex + 2, nil
			},
		}

		elp := NewExpressionListParser(fakeExpressionParser)
		expressionListNode, newIndex, err := elp.Parse(tokens, 0)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(expressionListNode, expectedExpressionListNode) {
			println(expressionListNode.Children[0].Token.Type)
			println(expectedExpressionListNode.Children[0].Token.Type)
			t.Errorf("Expected %v, got %v", expectedExpressionListNode, expressionListNode)
		}
		if newIndex != 2 {
			t.Errorf("Expected index 2, got %d", newIndex)
		}
	})
}

type fakeExpressionParser struct {
	ParseMock func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error)
}

func (f fakeExpressionParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	return f.ParseMock(tokens, currentIndex)
}
