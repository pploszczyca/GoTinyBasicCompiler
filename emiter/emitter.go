package emiter

import (
	"GoTinyBasicCompiler/domain"
	"strings"
)

type Emitter interface {
	Emit(programTree *domain.ProgramTree) (string, error)
}

type cEmitter struct {
	tokenEmitter TokenEmitter
}

func NewCEmitter(
	tokenEmitter TokenEmitter,
) Emitter {
	return &cEmitter{
		tokenEmitter: tokenEmitter,
	}
}

func (c *cEmitter) Emit(programTree *domain.ProgramTree) (string, error) {
	var builder strings.Builder

	builder.WriteString("#include <stdio.h>\n")
	builder.WriteString("int main() {\n")

	for _, node := range programTree.Nodes {
		c.emitNode(&builder, node, 1)
	}

	builder.WriteString("  return 0;\n")
	builder.WriteString("}\n")

	return builder.String(), nil
}

func (c *cEmitter) emitNode(builder *strings.Builder, node *domain.Node, indent int) {

}
