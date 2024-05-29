package domain

import "testing"

func TestTokenIterator_Current(t *testing.T) {
	t.Run("returns current token", func(t *testing.T) {
		tokens := []Token{
			{Type: Number},
			{Type: Identifier},
		}
		iterator := NewTokenIterator(tokens)

		currentToken, err := iterator.Current()

		if currentToken != tokens[0] {
			t.Errorf("Expected %v, got %v", tokens[0], currentToken)
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("returns error when index is out of range", func(t *testing.T) {
		iterator := NewTokenIterator([]Token{})

		_, err := iterator.Current()

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}

func TestTokenIterator_Next(t *testing.T) {
	t.Run("returns next token", func(t *testing.T) {
		tokens := []Token{
			{Type: Number},
			{Type: Identifier},
		}
		iterator := NewTokenIterator(tokens)

		iterator.Next()
		currentToken, err := iterator.Current()

		if currentToken != tokens[1] {
			t.Errorf("Expected %v, got %v", tokens[1], currentToken)
		}
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func TestTokenIterator_HasNext(t *testing.T) {
	t.Run("returns true when there is next token", func(t *testing.T) {
		iterator := NewTokenIterator([]Token{{Type: Number}})

		hasNext := iterator.HasNext()

		if !hasNext {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("returns false when there is no next token", func(t *testing.T) {
		iterator := NewTokenIterator([]Token{})

		hasNext := iterator.HasNext()

		if hasNext {
			t.Errorf("Expected false, got true")
		}
	})
}
