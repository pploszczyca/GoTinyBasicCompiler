package varList

import (
	"GoTinyBasicCompiler/domain"
	"reflect"
	"testing"
)

func TestVarListParser_Parse(t *testing.T) {
	t.Run("returns error when token is out of range", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{})
		expectedError := "tokens index out of range"

		vp := NewVarListParser()
		_, err := vp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when token is not an identifier", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Number}})
		expectedError := "expected identifier"

		vp := NewVarListParser()
		_, err := vp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when index is out of range after identifier", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier, Value: "A"}})
		expectedError := "tokens index out of range"

		vp := NewVarListParser()
		_, err := vp.Parse(&iterator)

		if err.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("returns correct node", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier, Value: "A"}, {Type: domain.Cr}})
		expectedVarListNode := domain.Node{Type: domain.VarListNode}
		expectedVarListNode.AddChildToken(domain.Token{Type: domain.Identifier, Value: "A"})

		vp := NewVarListParser()
		varListNode, err := vp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(varListNode, &expectedVarListNode) {
			t.Errorf("Expected %v, got %v", expectedVarListNode, varListNode)
		}
	})

	t.Run("returns correct node with multiple identifiers", func(t *testing.T) {
		iterator := domain.NewTokenIterator([]domain.Token{
			{Type: domain.Identifier, Value: "A"},
			{Type: domain.Comma},
			{Type: domain.Identifier, Value: "B"},
			{Type: domain.Comma},
			{Type: domain.Identifier, Value: "C"},
			{Type: domain.Cr},
		})
		expectedVarListNode := domain.Node{
			Type: domain.VarListNode,
			Children: []*domain.Node{
				{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
				{Token: domain.Token{Type: domain.Comma}},
				{Token: domain.Token{Type: domain.Identifier, Value: "B"}},
				{Token: domain.Token{Type: domain.Comma}},
				{Token: domain.Token{Type: domain.Identifier, Value: "C"}},
			},
		}
		vp := NewVarListParser()
		varListNode, err := vp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(varListNode, &expectedVarListNode) {
			t.Errorf("Expected %v, got %v", expectedVarListNode, varListNode)
		}
	})
}
