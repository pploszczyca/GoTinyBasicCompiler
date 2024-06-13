package emiter

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"testing"
)

func TestCEmitter_Emit(t *testing.T) {
	t.Run("returns C empty program when program tree is empty", func(t *testing.T) {
		programTree := &domain.ProgramTree{}
		expectedResult := `#include <stdio.h>
int main() {
}
`
		emitter := NewCEmitter(&testutils.FakeTokenEmitter{})
		result, err := emitter.Emit(programTree)

		if err != nil {
			t.Errorf("Error: %s\n", err)
		}
		if result != expectedResult {
			t.Errorf("Expected: %s, but got: %s\n", expectedResult, result)
		}
	})

	t.Run("returns C program with end token", func(t *testing.T) {
		endToken := domain.Token{Type: domain.End}
		lineNode := &domain.Node{
			Type: domain.LineNode,
			Children: []*domain.Node{
				{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
				{Type: domain.IdentifierNode, Token: endToken},
			},
		}
		programTree := &domain.ProgramTree{Nodes: []*domain.Node{lineNode}}
		expectedResult := `#include <stdio.h>
int main() {
    label_10:
    return 0;
}
`

		emitter := NewCEmitter(NewCTokenEmitter())
		result, err := emitter.Emit(programTree)

		if err != nil {
			t.Errorf("Error: %s\n", err)
		}
		if result != expectedResult {
			t.Errorf("Expected: \n%s, but got: \n%s\n", expectedResult, result)
		}
	})

	t.Run("returns C program with print node", func(t *testing.T) {
		lineNode := &domain.Node{
			Type: domain.LineNode,
			Children: []*domain.Node{
				{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
				{
					Type: domain.StatementNode,
					Children: []*domain.Node{
						{Token: domain.Token{Type: domain.Print}},
						{
							Type: domain.ExpressionListNode,
							Children: []*domain.Node{
								{Token: domain.Token{Type: domain.String, Value: "Hello"}},
								{Token: domain.Token{Type: domain.Comma}},
								{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
								{Token: domain.Token{Type: domain.Comma}},
								{
									Type: domain.ExpressionNode,
									Children: []*domain.Node{
										{Token: domain.Token{Type: domain.Number, Value: "1"}},
										{Token: domain.Token{Type: domain.Plus}},
										{Token: domain.Token{Type: domain.Number, Value: "2"}},
									},
								},
							},
						},
					},
				},
			},
		}
		programTree := &domain.ProgramTree{Nodes: []*domain.Node{lineNode}}
		expectedResult := `#include <stdio.h>
int main() {
    label_10:
    printf("%s,%d,%d", "Hello", A, 1+2);
}
`

		emitter := NewCEmitter(NewCTokenEmitter())
		result, err := emitter.Emit(programTree)

		if err != nil {
			t.Errorf("Error: %s\n", err)
		}
		if result != expectedResult {
			t.Errorf("Expected: \n%s, but got: \n%s\n", expectedResult, result)
		}
	})
}
