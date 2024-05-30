package domain

import (
	"reflect"
	"testing"
)

func TestNode_AddChildToken(t *testing.T) {
	t.Run("adds token to node", func(t *testing.T) {
		node := Node{}
		token := Token{Type: Identifier, Value: "A"}
		nodeToken := Node{Token: token}
		node.AddChildToken(token)

		if !reflect.DeepEqual(node.Children[0], &nodeToken) {
			t.Errorf("expected %v but got %v", nodeToken, node.Children[0])
		}
	})

}

func TestNode_AddChild(t *testing.T) {
	t.Run("adds child to node", func(t *testing.T) {
		node := Node{}
		child := &Node{}
		secondChild := &Node{}
		expectedChildren := []*Node{child, secondChild}

		node.AddChild(child)
		node.AddChild(secondChild)

		if !reflect.DeepEqual(node.Children, expectedChildren) {
			t.Errorf("expected %v but got %v", expectedChildren, node.Children)
		}
	})
}

func TestNode_IsLeaf(t *testing.T) {
	t.Run("returns false if node doesnt have children and token", func(t *testing.T) {
		node := Node{}
		if node.IsLeaf() {
			t.Errorf("expected false but got true")
		}
	})

	t.Run("returns false if node has children", func(t *testing.T) {
		node := Node{}
		node.AddChild(&Node{})
		if node.IsLeaf() {
			t.Errorf("expected false but got true")
		}
	})

	t.Run("returns true if node has token", func(t *testing.T) {
		node := Node{Token: Token{Type: Identifier, Value: "A"}}
		if !node.IsLeaf() {
			t.Errorf("expected true but got false")
		}
	})
}
