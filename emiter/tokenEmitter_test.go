package emiter

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"go/token"
	"testing"
)

func TestCTokenEmitter_Emit(t *testing.T) {
	tests := []struct {
		token         domain.Token
		expectedValue string
	}{
		{domain.Token{Type: domain.Number, Value: "1"}, "1"},
		{domain.Token{Type: domain.Identifier, Value: "A"}, "A"},
		{domain.Token{Type: domain.String, Value: "Hello"}, "Hello"},
		{domain.Token{Type: domain.Print}, "printf"},
		{domain.Token{Type: domain.Input}, "scanf"},
		{domain.Token{Type: domain.Plus}, "+"},
		{domain.Token{Type: domain.Minus}, "-"},
		{domain.Token{Type: domain.Multiply}, "*"},
		{domain.Token{Type: domain.Divide}, "/"},
		{domain.Token{Type: domain.Equal}, "="},
		{domain.Token{Type: domain.LessThan}, "<"},
		{domain.Token{Type: domain.GreaterThan}, ">"},
		{domain.Token{Type: domain.LessThanOrEqual}, "<="},
		{domain.Token{Type: domain.GreaterThanOrEqual}, ">="},
		{domain.Token{Type: domain.NotEqual}, "!="},
		{domain.Token{Type: domain.Comma}, ", "},
		{domain.Token{Type: domain.Semicolon}, ";"},
		{domain.Token{Type: domain.LParen}, "("},
		{domain.Token{Type: domain.RParen}, ")"},
		{domain.Token{Type: domain.End}, "return 0"},
		{domain.Token{Type: domain.Let}, "int"},
		{domain.Token{Type: domain.If}, "if"},
		{domain.Token{Type: domain.Goto}, "goto"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("returns C token for %s", token.TYPE), func(t *testing.T) {
			emitter := NewCTokenEmitter()
			result, err := emitter.Emit(test.token)
			if err != nil {
				t.Errorf("Error: %s\n", err)
			}
			if result != test.expectedValue {
				t.Errorf("Expected: %s, but got: %s\n", test.expectedValue, result)
			}
		})
	}

	t.Run("returns error for unknown token type", func(t *testing.T) {
		unknownToken := domain.Token{Type: domain.TokenType(999)}
		expectedError := fmt.Sprintf("Unknown token type: %s\n", unknownToken.Type)

		emitter := NewCTokenEmitter()
		_, err := emitter.Emit(unknownToken)

		if err.Error() != expectedError {
			t.Errorf("Expected: %s, but got: %s\n", expectedError, err)
		}
	})
}
