package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"fmt"
	"reflect"
	"testing"
)

func TestStatementParser_Parse(t *testing.T) {
	t.Run("parses print statement", func(t *testing.T) {
		expressionListNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionListNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(&fakeExpressionListParser)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("returns error when expression list parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(&fakeExpressionListParser)
		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("return error when unexpected statement", func(t *testing.T) {
		tokens := []domain.Token{
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("unexpected statement: 2")

		sp := NewStatementParser(&testutils.FakeNodeParser{})
		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})
}
