package statement

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/testutils"
	"fmt"
	"reflect"
	"testing"
)

func TestStatementParser_Parse(t *testing.T) {
	t.Run("parses print statement", func(t *testing.T) {
		expressionListNode := &domain.Node{Type: domain.ExpressionListNode}
		fakeExpressionListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionListNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionListNode},
			},
		}

		sp := NewStatementParser(
			&fakeExpressionListParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("returns error when expression list parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Print},
			{Type: domain.String},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&fakeExpressionListParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("return error when unexpected statement", func(t *testing.T) {
		tokens := []domain.Token{
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("unexpected statement: Identifier")

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			&fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and relop parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and second expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		index := 0
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				if index == 0 {
					index++
					return &domain.Node{Type: domain.ExpressionNode}, nil
				}

				return nil, expectedError
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.RelopNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and expression are parsed and iterator is out of index", func(t *testing.T) {
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.RelopNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("tokens index out of range")

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and expression are parsed and expected THEN token is not found", func(t *testing.T) {
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.RelopNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
			{Type: domain.Identifier},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("expected THEN")

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and expression are parsed and THEN is found and statement parsing return errro", func(t *testing.T) {
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.RelopNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
			{Type: domain.Identifier},
			{Type: domain.Then},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("returns correctly parsed if statement", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		relopNode := &domain.Node{Type: domain.RelopNode}
		expressionListNode := &domain.Node{Type: domain.ExpressionListNode}
		fakeExpressionListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionListNode, nil
			},
		}
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return relopNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.If},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
			{Type: domain.Identifier},
			{Type: domain.Then},
			{Type: domain.Print},
			{Type: domain.String},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			fakeExpressionListParser,
			fakeExpressionParser,
			fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if statementNode.Type != domain.StatementNode {
			t.Errorf("Expected StatementNode, got %v", statementNode.Type)
		}
		if len(statementNode.Children) != 6 {
			t.Errorf("Expected 6 children, got %v", len(statementNode.Children))
		}
		if statementNode.Children[0].Token.Type != domain.If {
			t.Errorf("Expected first child to be If token, got %v", statementNode.Children[0].Token.Type)
		}
		if statementNode.Children[1].Type != domain.ExpressionNode {
			t.Errorf("Expected second child to be ExpressionNode, got %v", statementNode.Children[1].Type)
		}
		if statementNode.Children[2].Type != domain.RelopNode {
			t.Errorf("Expected third child to be RelopNode, got %v", statementNode.Children[2].Type)
		}
		if statementNode.Children[3].Type != domain.ExpressionNode {
			t.Errorf("Expected fourth child to be ExpressionNode, got %v", statementNode.Children[3].Type)
		}
		if statementNode.Children[4].Token.Type != domain.Then {
			t.Errorf("Expected fifth child to be Then token, got %v", statementNode.Children[4].Token.Type)
		}
		if statementNode.Children[5].Type != domain.StatementNode {
			t.Errorf("Expected sixth child to be StatementNode, got %v", statementNode.Children[5].Type)
		}
		if statementNode.Children[5].Children[0].Token.Type != domain.Print {
			t.Errorf("Expected sixth child to be Print token, got %v", statementNode.Children[5].Children[0].Token.Type)
		}
		if statementNode.Children[5].Children[1].Type != domain.ExpressionListNode {
			t.Errorf("Expected sixth child to be ExpressionListNode, got %v", statementNode.Children[5].Children[1].Type)
		}
	})

	t.Run("returns error when goto token and expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Goto},
			{Type: domain.String},
			{Type: domain.Cr},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns successful goto node", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Goto},
			{Type: domain.String},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("returns error when var list parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeVarListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Input},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&fakeVarListParser,
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("parses input statement successfully", func(t *testing.T) {
		varListNode := &domain.Node{Type: domain.VarListNode}
		fakeVarListParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return varListNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Input},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.VarListNode},
			},
		}

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
			&fakeVarListParser,
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("parses let statement successfully", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Let},
			{Type: domain.Identifier},
			{Type: domain.Equal},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Token: tokens[1]},
				{Token: tokens[2]},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("returns error when expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Let},
			{Type: domain.Identifier},
			{Type: domain.Equal},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when expected equal token is not found", func(t *testing.T) {
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Let},
			{Type: domain.Identifier},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("expected equal")

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when expected identifier token is not found", func(t *testing.T) {
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Let},
			{Type: domain.Equal},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedError := fmt.Errorf("expected identifier")

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("parses gosub statement successfully", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.Gosub},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	t.Run("returns error when expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.Gosub},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and expression parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.While},
			{Type: domain.Identifier},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			&fakeExpressionParser,
			&testutils.FakeNodeParser{},
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("returns error when if token and relop parser returns error", func(t *testing.T) {
		expectedError := fmt.Errorf("parse error")
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.ExpressionNode}, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				return nil, expectedError
			},
		}
		tokens := []domain.Token{
			{Type: domain.While},
			{Type: domain.Identifier},
			{Type: domain.LessThan},
		}
		iterator := domain.NewTokenIterator(tokens)

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			&fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		_, err := sp.Parse(&iterator)

		if !reflect.DeepEqual(err, expectedError) {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	})

	t.Run("parse while statement successfully", func(t *testing.T) {
		expressionNode := &domain.Node{Type: domain.ExpressionNode}
		fakeExpressionParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return expressionNode, nil
			},
		}
		fakeRelopParser := testutils.FakeNodeParser{
			ParseMock: func(iterator *domain.TokenIterator) (*domain.Node, error) {
				iterator.Next()
				return &domain.Node{Type: domain.RelopNode}, nil
			},
		}
		tokens := []domain.Token{
			{Type: domain.While},
			{Type: domain.Identifier, Value: "A"},
			{Type: domain.LessThan},
			{Type: domain.Number, Value: "10"},
		}
		iterator := domain.NewTokenIterator(tokens)
		expectedStatementNode := &domain.Node{
			Type: domain.StatementNode,
			Children: []*domain.Node{
				{Token: tokens[0]},
				{Type: domain.ExpressionNode},
				{Type: domain.RelopNode},
				{Type: domain.ExpressionNode},
			},
		}

		sp := NewStatementParser(
			&testutils.FakeNodeParser{},
			fakeExpressionParser,
			fakeRelopParser,
			&testutils.FakeNodeParser{},
		)

		statementNode, err := sp.Parse(&iterator)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(statementNode, expectedStatementNode) {
			t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
		}
	})

	testCases := []struct {
		token domain.Token
	}{
		{token: domain.Token{Type: domain.Return}},
		{token: domain.Token{Type: domain.Clear}},
		{token: domain.Token{Type: domain.List}},
		{token: domain.Token{Type: domain.Run}},
		{token: domain.Token{Type: domain.End}},
		{token: domain.Token{Type: domain.Wend}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("parses %v statement successfully", tc.token.Type), func(t *testing.T) {
			tokens := []domain.Token{
				tc.token,
			}
			iterator := domain.NewTokenIterator(tokens)
			expectedStatementNode := &domain.Node{
				Type: domain.StatementNode,
				Children: []*domain.Node{
					{Token: tokens[0]},
				},
			}

			sp := NewStatementParser(
				&testutils.FakeNodeParser{},
				&testutils.FakeNodeParser{},
				&testutils.FakeNodeParser{},
				&testutils.FakeNodeParser{},
			)

			statementNode, err := sp.Parse(&iterator)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(statementNode, expectedStatementNode) {
				t.Errorf("Expected %v, got %v", expectedStatementNode, statementNode)
			}
		})

	}
}
