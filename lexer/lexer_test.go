package lexer

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	lexer := NewLexer()

	t.Run("returns tokens without error", func(t *testing.T) {
		inputCode := "10 PRINT \"HELLO, WORLD!\"\n20 GOTO 10\n"
		expectedTokens := []domain.Token{
			{Type: domain.Number, Value: "10"},
			{Type: domain.Print},
			{Type: domain.String, Value: "\"HELLO, WORLD!\""},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "20"},
			{Type: domain.Goto},
			{Type: domain.Number, Value: "10"},
			{Type: domain.Cr},
			{Type: domain.Eof},
		}

		tokens, err := lexer.Lex(inputCode)

		if err != nil {
			t.Errorf("error should be nil, but got %v", err)
		}

		if !reflect.DeepEqual(tokens, expectedTokens) {
			t.Errorf("expected %v, but got %v", expectedTokens, tokens)
		}
	})

	t.Run("returns tokens for more complicated code", func(t *testing.T) {
		inputCode := `10 PRINT "ENTER FIRST NUMBER: "
20 INPUT A
30 PRINT "ENTER SECOND NUMBER: "
40 INPUT B
50 LET C = A + B
60 PRINT "THE SUM IS: "; C
70 END`
		expectedTokens := []domain.Token{
			{Type: domain.Number, Value: "10"},
			{Type: domain.Print},
			{Type: domain.String, Value: "\"ENTER FIRST NUMBER: \""},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "20"},
			{Type: domain.Input},
			{Type: domain.Identifier, Value: "A"},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "30"},
			{Type: domain.Print},
			{Type: domain.String, Value: "\"ENTER SECOND NUMBER: \""},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "40"},
			{Type: domain.Input},
			{Type: domain.Identifier, Value: "B"},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "50"},
			{Type: domain.Let},
			{Type: domain.Identifier, Value: "C"},
			{Type: domain.Equal},
			{Type: domain.Identifier, Value: "A"},
			{Type: domain.Plus},
			{Type: domain.Identifier, Value: "B"},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "60"},
			{Type: domain.Print},
			{Type: domain.String, Value: "\"THE SUM IS: \""},
			{Type: domain.Semicolon},
			{Type: domain.Identifier, Value: "C"},
			{Type: domain.Cr},
			{Type: domain.Number, Value: "70"},
			{Type: domain.End},
			{Type: domain.Eof},
		}

		tokens, err := lexer.Lex(inputCode)

		if err != nil {
			t.Errorf("error should be nil, but got %v", err)
		}

		if !reflect.DeepEqual(tokens, expectedTokens) {
			t.Errorf("expected %v, but got %v", expectedTokens, tokens)
		}
	})

	keywordsTestCases := []struct {
		inputCode     string
		expectedToken domain.Token
	}{
		{
			inputCode:     "PRINT",
			expectedToken: domain.Token{Type: domain.Print},
		},
		{
			inputCode:     "IF",
			expectedToken: domain.Token{Type: domain.If},
		},
		{
			inputCode:     "THEN",
			expectedToken: domain.Token{Type: domain.Then},
		},
		{
			inputCode:     "GOTO",
			expectedToken: domain.Token{Type: domain.Goto},
		},
		{
			inputCode:     "INPUT",
			expectedToken: domain.Token{Type: domain.Input},
		},
		{
			inputCode:     "LET",
			expectedToken: domain.Token{Type: domain.Let},
		},
		{
			inputCode:     "GOSUB",
			expectedToken: domain.Token{Type: domain.Gosub},
		},
		{
			inputCode:     "RETURN",
			expectedToken: domain.Token{Type: domain.Return},
		},
		{
			inputCode:     "CLEAR",
			expectedToken: domain.Token{Type: domain.Clear},
		},
		{
			inputCode:     "LIST",
			expectedToken: domain.Token{Type: domain.List},
		},
		{
			inputCode:     "RUN",
			expectedToken: domain.Token{Type: domain.Run},
		},
		{
			inputCode:     "END",
			expectedToken: domain.Token{Type: domain.End},
		},
		{
			inputCode:     "WHEN",
			expectedToken: domain.Token{Type: domain.While},
		},
		{
			inputCode:     "WEND",
			expectedToken: domain.Token{Type: domain.Wend},
		},
		{
			inputCode:     "+",
			expectedToken: domain.Token{Type: domain.Plus},
		},
		{
			inputCode:     "-",
			expectedToken: domain.Token{Type: domain.Minus},
		},
		{
			inputCode:     "*",
			expectedToken: domain.Token{Type: domain.Multiply},
		},
		{
			inputCode:     "/",
			expectedToken: domain.Token{Type: domain.Divide},
		},
		{
			inputCode:     "=",
			expectedToken: domain.Token{Type: domain.Equal},
		},
		{
			inputCode:     "<",
			expectedToken: domain.Token{Type: domain.LessThan},
		},
		{
			inputCode:     ">",
			expectedToken: domain.Token{Type: domain.GreaterThan},
		},
		{
			inputCode:     "<=",
			expectedToken: domain.Token{Type: domain.LessThanOrEqual},
		},
		{
			inputCode:     ">=",
			expectedToken: domain.Token{Type: domain.GreaterThanOrEqual},
		},
		{
			inputCode:     "<>",
			expectedToken: domain.Token{Type: domain.NotEqual},
		},
		{
			inputCode:     ",",
			expectedToken: domain.Token{Type: domain.Comma},
		},
		{
			inputCode:     ";",
			expectedToken: domain.Token{Type: domain.Semicolon},
		},
		{
			inputCode:     "(",
			expectedToken: domain.Token{Type: domain.LParen},
		},
		{
			inputCode:     ")",
			expectedToken: domain.Token{Type: domain.RParen},
		},
		{
			inputCode:     "123",
			expectedToken: domain.Token{Type: domain.Number, Value: "123"},
		},
		{
			inputCode:     "\"abc\"",
			expectedToken: domain.Token{Type: domain.String, Value: "\"abc\""},
		},
		{
			inputCode:     "A",
			expectedToken: domain.Token{Type: domain.Identifier, Value: "A"},
		},
	}

	for _, tc := range keywordsTestCases {
		t.Run(fmt.Sprintf("Returns correct token for input: %s", tc.inputCode), func(t *testing.T) {
			expectedTokens := []domain.Token{
				tc.expectedToken,
				{Type: domain.Eof},
			}

			tokens, err := lexer.Lex(tc.inputCode)

			if err != nil {
				t.Errorf("error should be nil, but got %v", err)
			}

			if !reflect.DeepEqual(tokens, expectedTokens) {
				t.Errorf("expected %v, but got %v", expectedTokens, tokens)
			}
		})
	}

	t.Run("returns error for invalid token", func(t *testing.T) {
		inputCode := "abcd"
		expectedError := fmt.Errorf("error lexing line 1: invalid token: abcd")

		_, err := lexer.Lex(inputCode)

		if err == nil {
			t.Errorf("error should not be nil")
		}

		if err.Error() != expectedError.Error() {
			t.Errorf("expected %v, but got %v", expectedError, err)
		}
	})

	t.Run("returns correct tokens for print statement with multiple expressions", func(t *testing.T) {
		inputCode := "260 PRINT K, \" SQUARED IS \", Q\n"
		expectedTokens := []domain.Token{
			{Type: domain.Number, Value: "260"},
			{Type: domain.Print},
			{Type: domain.Identifier, Value: "K"},
			{Type: domain.Comma},
			{Type: domain.String, Value: "\" SQUARED IS \""},
			{Type: domain.Comma},
			{Type: domain.Identifier, Value: "Q"},
			{Type: domain.Cr},
			{Type: domain.Eof},
		}

		tokens, err := lexer.Lex(inputCode)

		if err != nil {
			t.Errorf("error should be nil, but got %v", err)
		}

		if !reflect.DeepEqual(tokens, expectedTokens) {
			t.Errorf("expected %v, but got %v", expectedTokens, tokens)
		}
	})
}
