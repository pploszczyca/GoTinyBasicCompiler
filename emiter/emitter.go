package emiter

import (
	"GoTinyBasicCompiler/domain"
	"GoTinyBasicCompiler/utils"
	"strconv"
	"strings"
)

type Emitter interface {
	Emit(programTree *domain.ProgramTree) (string, error)
}

type cEmitter struct {
	tokenEmitter TokenEmitter
	cUtils       string
}

func NewTestedCTokenEmitter(
	tokenEmitter TokenEmitter,
	cUtils string,
) Emitter {
	return &cEmitter{
		tokenEmitter: tokenEmitter,
		cUtils:       cUtils,
	}
}

func NewCEmitter(
	tokenEmitter TokenEmitter,
) Emitter {
	return &cEmitter{
		tokenEmitter: tokenEmitter,
		cUtils: `typedef struct {
	int lineNumber;
	void *labelAddr;
} LabelMap;

void* find_label(int lineNumber, LabelMap labels[], int numLabels) {
	for (int i = 0; i < numLabels; ++i) {
		if (labels[i].lineNumber == lineNumber) {
			return labels[i].labelAddr;
		}
	}
}

#define MAX 100

typedef struct {
    int top;
    void* items[MAX];
} Stack;

void initStack(Stack* s) {
    s->top = -1;
}

int isEmpty(Stack* s) {
    return s->top == -1;
}

int isFull(Stack* s) {
    return s->top == MAX - 1;
}

void push(Stack* s, void* label) {
    if (isFull(s)) {
        return;
    }
    s->items[++(s->top)] = label;
}

void* pop(Stack* s) {
    if (isEmpty(s)) {
        return NULL;
    }
    return s->items[(s->top)--];
}

void* peek(Stack* s) {
    if (isEmpty(s)) {
        return NULL;
    }
    return s->items[s->top];
}

`,
	}
}

func (c *cEmitter) Emit(programTree *domain.ProgramTree) (string, error) {
	var builder strings.Builder
	previousIdentifiers := utils.NewSet[domain.Token]()
	goSubIndex := 0

	builder.WriteString("#include <stdio.h>\n\n")
	builder.WriteString(c.cUtils)
	builder.WriteString(`int main() {
	Stack gosubStack;
	initStack(&gosubStack);
`)

	c.emitLabelsMap(&builder, programTree.Nodes, 1)

	err := c.emitMultipleNodes(&builder, programTree.Nodes, previousIdentifiers, 1, &goSubIndex)
	if err != nil {
		return "", err
	}

	builder.WriteString("}\n")

	return builder.String(), nil
}

func (c *cEmitter) emitLabelsMap(
	builder *strings.Builder,
	nodes []*domain.Node,
	indent int,
) {
	c.writeIndent(builder, indent)
	builder.WriteString("LabelMap labels[] = {\n")

	for _, node := range nodes {
		c.writeIndent(builder, indent+1)
		if node.Type == domain.LineNode && node.Children[0].Type == domain.NumberNode {
			lineValue := node.Children[0].Token.Value
			builder.WriteString("{" + lineValue + ", &&label_" + lineValue + "},\n")
		}
	}

	c.writeIndent(builder, indent)
	builder.WriteString("};\n")
	c.writeIndent(builder, indent)
	builder.WriteString("int numLabels = sizeof(labels) / sizeof(labels[0]);\n")
}

func (c *cEmitter) emitMultipleNodes(
	builder *strings.Builder,
	nodes []*domain.Node,
	previousIdentifiers *utils.Set[domain.Token],
	indent int,
	goSubIndex *int,
) error {
	for _, child := range nodes {
		err := c.emitNode(builder, child, previousIdentifiers, indent, goSubIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *cEmitter) emitNode(
	builder *strings.Builder,
	node *domain.Node,
	previousIdentifiers *utils.Set[domain.Token],
	indent int,
	goSubIndex *int,
) error {
	if node.IsLeaf() {
		stringToken, err := c.tokenEmitter.Emit(node.Token)
		if err != nil {
			return err
		}

		if node.Token.Type == domain.Identifier {
			previousIdentifiers.Add(node.Token)
		}

		builder.WriteString(stringToken)
		return nil
	}

	switch node.Type {
	case domain.LineNode:
		return c.emitLineNode(builder, node, previousIdentifiers, indent, goSubIndex)
	case domain.ProgramNode:
	case domain.StatementNode:
		return c.emitStatementNode(builder, node, previousIdentifiers, indent, goSubIndex)
	case domain.ExpressionNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	case domain.NumberNode:
	case domain.IdentifierNode:
	case domain.OperatorNode:
	case domain.ExpressionListNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	case domain.TermNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	case domain.FactorNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	case domain.RelopNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	case domain.VarListNode:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	}

	return nil
}

func (c *cEmitter) emitLineNode(
	builder *strings.Builder,
	node *domain.Node,
	previousIdentifiers *utils.Set[domain.Token],
	indent int,
	goSubIndex *int,
) error {
	statementIndex := 0
	c.writeIndent(builder, indent)
	if node.Children[0].Type == domain.NumberNode {
		builder.WriteString("label_" + node.Children[0].Token.Value + ":\n")
		statementIndex = 1
		c.writeIndent(builder, indent)
	}

	err := c.emitMultipleNodes(builder, node.Children[statementIndex:], previousIdentifiers, indent, goSubIndex)
	if err != nil {
		return err
	}

	if c.shouldWriteEndLine(node) {
		builder.WriteString(";")
	}
	builder.WriteString("\n")
	return nil
}

func (c *cEmitter) emitStatementNode(
	builder *strings.Builder,
	node *domain.Node,
	previousIdentifiers *utils.Set[domain.Token],
	indent int,
	goSubIndex *int,
) error {
	switch node.Children[0].Token.Type {
	case domain.Print:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken + "(\"")
		expressionListNode := node.Children[1]

		for _, child := range expressionListNode.Children {
			if child.Token.Type == domain.String {
				builder.WriteString("%s")
			} else if child.Token.Type != domain.Comma {
				builder.WriteString("%d")
			}
		}

		builder.WriteString("\\n\", ")

		err = c.emitNode(builder, expressionListNode, previousIdentifiers, indent, goSubIndex)
		if err != nil {
			return err
		}

		builder.WriteString(")")

	case domain.Let:
		childIndex := 0

		if node.Children[1].Token.Type == domain.Identifier && previousIdentifiers.Contains(node.Children[1].Token) {
			childIndex = 1
		}

		children := node.Children[childIndex:]
		for index, child := range children {
			err := c.emitNode(builder, child, previousIdentifiers, indent, goSubIndex)
			if err != nil {
				return err
			}
			if index != len(children)-1 {
				builder.WriteString(" ")
			}
		}

	case domain.If:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken + " (")

		for _, child := range node.Children[1:] {
			if child.Token.Type == domain.Then {
				builder.WriteString(") ")
			} else {
				err := c.emitNode(builder, child, previousIdentifiers, indent, goSubIndex)
				if err != nil {
					return err
				}
			}
		}

	case domain.Goto:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken + " *find_label(")

		if err := c.emitNode(builder, node.Children[1], previousIdentifiers, indent, goSubIndex); err != nil {
			return err
		}

		builder.WriteString(", labels, numLabels)")

		return nil

	case domain.Input:
		builder.WriteString("int ")
		inputNode := node.Children[0]
		varListNode := node.Children[1]

		err := c.emitMultipleNodes(builder, varListNode.Children, previousIdentifiers, indent, goSubIndex)
		if err != nil {
			return err
		}

		builder.WriteString(";\n")
		c.writeIndent(builder, indent)

		stringToken, err := c.tokenEmitter.Emit(inputNode.Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken + "(\"")

		for _, child := range varListNode.Children {
			if child.Token.Type == domain.Comma {
				builder.WriteString(",")
			} else {
				builder.WriteString("%d")
			}
		}

		builder.WriteString("\", ")

		for _, child := range varListNode.Children {
			if child.Token.Type != domain.Comma {
				builder.WriteString("&")
			}
			err := c.emitNode(builder, child, previousIdentifiers, indent, goSubIndex)
			if err != nil {
				return err
			}
		}
		builder.WriteString(")")

	case domain.While:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}

		builder.WriteString(stringToken + " (")
		err = c.emitMultipleNodes(builder, node.Children[1:], previousIdentifiers, indent, goSubIndex)
		if err != nil {
			return err
		}

		builder.WriteString(") {")

	case domain.For:
		forToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}

		identifierString, err := c.tokenEmitter.Emit(node.Children[1].Token)
		if err != nil {
			return err
		}

		equalString, err := c.tokenEmitter.Emit(node.Children[2].Token)
		if err != nil {
			return err
		}

		var fromStringBuilder strings.Builder
		if err := c.emitNode(&fromStringBuilder, node.Children[3], previousIdentifiers, indent, goSubIndex); err != nil {
			return err
		}

		_, err = c.tokenEmitter.Emit(node.Children[4].Token)
		if err != nil {
			return err
		}

		var toStringBuilder strings.Builder
		if err := c.emitNode(&toStringBuilder, node.Children[5], previousIdentifiers, indent, goSubIndex); err != nil {
			return err
		}

		fromExpression := fromStringBuilder.String()
		toExpression := toStringBuilder.String()

		builder.WriteString(forToken + " (int " + identifierString + " " + equalString + " " + fromExpression + "; " + identifierString + " <= " + toExpression + "; " + identifierString + "++) {")

	case domain.Next:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		builder.WriteString(stringToken)

	case domain.Gosub:
		stringToken, err := c.tokenEmitter.Emit(node.Children[0].Token)
		if err != nil {
			return err
		}
		*goSubIndex = *goSubIndex + 1
		builder.WriteString("push(&gosubStack, &&label_gosub_" + strconv.Itoa(*goSubIndex) + ");\n")
		c.writeIndent(builder, indent)

		builder.WriteString(stringToken + " *find_label(")

		if err := c.emitNode(builder, node.Children[1], previousIdentifiers, indent, goSubIndex); err != nil {
			return err
		}

		builder.WriteString(", labels, numLabels);\n")
		c.writeIndent(builder, indent)
		builder.WriteString("label_gosub_" + strconv.Itoa(*goSubIndex) + ":")

	case domain.Return:
		builder.WriteString("goto *pop(&gosubStack)")

	default:
		return c.emitMultipleNodes(builder, node.Children, previousIdentifiers, indent, goSubIndex)
	}

	return nil
}

func (c *cEmitter) shouldWriteEndLine(node *domain.Node) bool {
	if node.IsLeaf() {
		tokenType := node.Token.Type
		return tokenType != domain.While && tokenType != domain.Wend && tokenType != domain.Next && tokenType != domain.For && tokenType != domain.Gosub
	}

	for _, child := range node.Children {
		if !c.shouldWriteEndLine(child) {
			return false
		}
	}
	return true
}

func (c *cEmitter) writeIndent(builder *strings.Builder, indent int) {
	builder.WriteString(strings.Repeat("\t", indent))
}
