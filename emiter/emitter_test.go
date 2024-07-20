package emiter

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"errors"
	"fmt"
	"testing"
)

func TestCEmitter_Emit(t *testing.T) {
	tests := []struct {
		name           string
		programTree    *domain.ProgramTree
		expectedResult string
		tokenEmitter   TokenEmitter
	}{
		{
			name:        "returns C empty program when program tree is empty",
			programTree: &domain.ProgramTree{},
			expectedResult: `#include <stdio.h>
int main() {
}
`,
			tokenEmitter: &testutils.FakeTokenEmitter{},
		},
		{
			name: "returns C program with end token",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{Type: domain.StatementNode, Children: []*domain.Node{{Token: domain.Token{Type: domain.End}}}},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    return 0;
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with print node",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
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
											{Token: domain.Token{Type: domain.String, Value: "\"Hello\""}},
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
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    printf("%s%d%d\n", "Hello", A, 1+2);
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with if statement",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.If}},
									{
										Type: domain.ExpressionNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
										},
									},
									{
										Type: domain.RelopNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.LessThan}},
										},
									},
									{
										Type: domain.ExpressionNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Number, Value: "5"}},
										},
									},
									{Token: domain.Token{Type: domain.Then}},
									{
										Type: domain.StatementNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Print}},
											{
												Type: domain.ExpressionListNode,
												Children: []*domain.Node{
													{Token: domain.Token{Type: domain.String, Value: "\"Hello\""}},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    if (A<5) printf("%s\n", "Hello");
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with goto statement",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Goto}},
									{
										Type: domain.ExpressionNode,
										Children: []*domain.Node{
											{
												Type: domain.TermNode,
												Children: []*domain.Node{
													{
														Type: domain.FactorNode,
														Children: []*domain.Node{
															{Token: domain.Token{Type: domain.Number, Value: "20"}},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    goto label_20;
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with input statement",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Input}},
									{
										Type: domain.VarListNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
											{Token: domain.Token{Type: domain.Comma}},
											{Token: domain.Token{Type: domain.Identifier, Value: "B"}},
										},
									},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    int A, B;
    scanf("%d,%d", &A, &B);
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with setting multiple times value to the identifier",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Let}},
									{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
									{Token: domain.Token{Type: domain.Equal}},
									{Token: domain.Token{Type: domain.Number, Value: "1"}},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "20"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Let}},
									{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
									{Token: domain.Token{Type: domain.Equal}},
									{Token: domain.Token{Type: domain.Number, Value: "2"}},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    int A = 1;
    label_20:
    A = 2;
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with while statement",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.While}},
									{
										Type: domain.ExpressionNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
										},
									},
									{
										Type: domain.RelopNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.LessThan}},
										},
									},
									{
										Type: domain.ExpressionNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.Number, Value: "5"}},
										},
									},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "20"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Print}},
									{
										Type: domain.ExpressionListNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.String, Value: "\"Hello\""}},
										},
									},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "30"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Let}},
									{Token: domain.Token{Type: domain.Identifier, Value: "A"}},
									{Token: domain.Token{Type: domain.Equal}},
									{Token: domain.Token{Type: domain.Number, Value: "1"}},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "40"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Wend}},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    while (A<5) {
    label_20:
    printf("%s\n", "Hello");
    label_30:
    A = 1;
    label_40:
    }
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
		{
			name: "returns C program with for statement",
			programTree: &domain.ProgramTree{
				Nodes: []*domain.Node{
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "10"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.For}},
									{Token: domain.Token{Type: domain.Identifier, Value: "I"}},
									{Token: domain.Token{Type: domain.Equal}},
									{Type: domain.ExpressionNode, Children: []*domain.Node{{Token: domain.Token{Type: domain.Number, Value: "1"}}}},
									{Token: domain.Token{Type: domain.To}},
									{Type: domain.ExpressionNode, Children: []*domain.Node{{Token: domain.Token{Type: domain.Number, Value: "5"}}}},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "20"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Print}},
									{
										Type: domain.ExpressionListNode,
										Children: []*domain.Node{
											{Token: domain.Token{Type: domain.String, Value: "\"Hello\""}},
										},
									},
								},
							},
						},
					},
					{
						Type: domain.LineNode,
						Children: []*domain.Node{
							{Type: domain.NumberNode, Token: domain.Token{Type: domain.Number, Value: "30"}},
							{
								Type: domain.StatementNode,
								Children: []*domain.Node{
									{Token: domain.Token{Type: domain.Next}},
									{Token: domain.Token{Type: domain.Identifier, Value: "I"}},
								},
							},
						},
					},
				},
			},
			expectedResult: `#include <stdio.h>
int main() {
    label_10:
    for (int I = 1; I <= 5; I++) {
    label_20:
    printf("%s\n", "Hello");
    label_30:
    }
}
`,
			tokenEmitter: NewCTokenEmitter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emitter := NewCEmitter(tt.tokenEmitter)
			result, err := emitter.Emit(tt.programTree)

			if err != nil {
				t.Errorf("Error: %s\n", err)
			}
			if result != tt.expectedResult {
				t.Errorf("Expected: \n%s, but got: \n%s\n", tt.expectedResult, result)
			}
		})
	}

	t.Run("returns error when token emitter fails", func(t *testing.T) {
		fakeError := fmt.Errorf("fake error")
		tokenEmitter := &testutils.FakeTokenEmitter{
			EmitMock: func(token domain.Token) (string, error) {
				return "", fakeError
			},
		}
		programTree := &domain.ProgramTree{
			Nodes: []*domain.Node{
				{
					Type: domain.LineNode,
					Children: []*domain.Node{
						{Token: domain.Token{Type: domain.Number, Value: "10"}},
					},
				},
			},
		}

		emitter := NewCEmitter(tokenEmitter)
		_, err := emitter.Emit(programTree)

		if !errors.Is(err, fakeError) {
			t.Errorf("Expected error: %s, but got: %s\n", fakeError, err)
		}
	})
}
