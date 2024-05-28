package statement

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"reflect"
	"testing"
)

func TestStatementParser_Parse(t *testing.T) {
	t.Run("parses print statement", func(t *testing.T) {
		expressionListNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionListParser := fakeExpressionListParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return expressionListNode, currentIndex + 1, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
		}
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(fakeExpressionListParser)
		statementNode, index, err := sp.Parse(tokens, 0)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
		if index != 2 {
			t.Errorf("Expected index 3, got %d", index)
		}
	})

	t.Run("returns error when expression list parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionListParser := fakeExpressionListParser{
			ParseMock: func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
				return nil, 0, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
			{Type: domain.Cr},
		}

		sp := NewStatementParser(fakeExpressionListParser)
		_, _, err := sp.Parse(tokens, 0)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("return error when unexpected statement", func(t *testing.T) {
		tokens := []domain.Token{
			{Type: domain.Identifier},
		}
		expectedError := fmt.Errorf("unexpected statement: 2")

		sp := NewStatementParser(fakeExpressionListParser{})
		_, _, err := sp.Parse(tokens, 0)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})
}

type fakeExpressionListParser struct {
	ParseMock func(tokens []domain.Token, currentIndex int) (*domain.Node, int, error)
}

func (f fakeExpressionListParser) Parse(tokens []domain.Token, currentIndex int) (*domain.Node, int, error) {
	return f.ParseMock(tokens, currentIndex)
}
