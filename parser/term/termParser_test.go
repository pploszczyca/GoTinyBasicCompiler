package term

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"fmt"
	"reflect"
	"testing"
)

func TestTermParser_Parse(t *testing.T) {
	t.Run("returns term node when factor parser returns node", func(t *testing.T) {
		factorNode := &domain.Node{Type: domain.FactorNode}
		fakeFactorParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return factorNode, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier}, {Type: domain.Cr}})
		expectedTermNode := &domain.Node{Type: domain.TermNode, Children: []*domain.Node{factorNode}}

		tp := NewTermParser(fakeFactorParser)
		termNode, err := tp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(termNode, expectedTermNode) {
			t.Errorf("Expected %v, got %v", expectedTermNode, termNode)
		}
	})

	t.Run("returns error when factor parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeFactorParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier}})

		tp := NewTermParser(fakeFactorParser)
		_, err := tp.Parse(&iterator)

		if err.Error() != expectedError.Error() {
			t.Errorf("Expected %v, got %v", expectedError, err)
		}
	})

	t.Run("return term node with multiple factor nodes", func(t *testing.T) {
		factorNode1 := &domain.Node{Type: domain.FactorNode}
		factorNode2 := &domain.Node{Type: domain.FactorNode}
		factorNode3 := &domain.Node{Type: domain.FactorNode}
		factorNodes := []*domain.Node{factorNode1, factorNode2, factorNode3}
		counter := 0

		fakeFactorParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				node := factorNodes[counter]
				counter++
				return node, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier}, {Type: domain.Multiply}, {Type: domain.Identifier}, {Type: domain.Divide}, {Type: domain.Identifier}, {Type: domain.Cr}})
		expectedTermNode := &domain.Node{Type: domain.TermNode, Children: []*domain.Node{
			factorNode1,
			{Token: domain.Token{Type: domain.Multiply}},
			factorNode2,
			{Token: domain.Token{Type: domain.Divide}},
			factorNode3,
		}}

		tp := NewTermParser(fakeFactorParser)
		termNode, err := tp.Parse(&iterator)

		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
		if !reflect.DeepEqual(termNode, expectedTermNode) {
			t.Errorf("Expected %v, got %v", expectedTermNode, termNode)
		}
	})

	t.Run("return error when index is out of range", func(t *testing.T) {
		factorNode1 := &domain.Node{Type: domain.FactorNode}
		factorNode2 := &domain.Node{Type: domain.FactorNode}
		factorNode3 := &domain.Node{Type: domain.FactorNode}
		factorNodes := []*domain.Node{factorNode1, factorNode2, factorNode3}
		counter := 0

		fakeFactorParser := &testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				node := factorNodes[counter]
				counter++
				return node, nil
			},
		}
		iterator := domain.NewTokenIterator([]domain.Token{{Type: domain.Identifier}, {Type: domain.Multiply}, {Type: domain.Identifier}, {Type: domain.Divide}, {Type: domain.Identifier}})

		tp := NewTermParser(fakeFactorParser)
		_, err := tp.Parse(&iterator)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
