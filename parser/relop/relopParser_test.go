package relop

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"reflect"
	"testing"
)

func TestRelopParser_Parse(t *testing.T) {
	allowedOperatorTestCases := []struct {
		tokenType domain.TokenType
	}{
		{domain.LessThan},
		{domain.LessThanOrEqual},
		{domain.GreaterThan},
		{domain.GreaterThanOrEqual},
		{domain.Equal},
		{domain.NotEqual},
	}

	for _, tc := range allowedOperatorTestCases {
		t.Run(fmt.Sprintf("returns correct node for token type: %d", tc.tokenType), func(t *testing.T) {
			iterator := domain.NewTokenIterator([]domain.Token{{Type: tc.tokenType}})
			tokenNode := domain.Node{Token: domain.Token{Type: tc.tokenType}}
			expectedNode := &domain.Node{Type: domain.RelopNode, Children: []*domain.Node{&tokenNode}}

			relopParser := NewRelopParser()
			node, err := relopParser.Parse(&iterator)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(node, expectedNode) {
				t.Errorf("expected node: %v, got: %v", expectedNode, node)
			}
		})
	}

	t.Run("returns error for unexpected token", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Plus}})
		relopParser := NewRelopParser()

		_, err := relopParser.Parse(&iterator)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
