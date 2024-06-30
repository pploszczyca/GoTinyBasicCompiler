package utils

import "testing"

func TestSet_Add(t *testing.T) {
	t.Run("adds element to set", func(t *testing.T) {
		set := NewSet[int]()
		set.Add(1)
		if !set.Contains(1) {
			t.Errorf("Expected set to contain 1")
		}
	})
}

func TestSet_Contains(t *testing.T) {
	t.Run("returns true when element is in set", func(t *testing.T) {
		set := NewSet[int]()
		set.Add(1)
		if !set.Contains(1) {
			t.Errorf("Expected set to contain 1")
		}
	})

	t.Run("returns false when element is not in set", func(t *testing.T) {
		set := NewSet[int]()
		if set.Contains(1) {
			t.Errorf("Expected set to not contain 1")
		}
	})
}
