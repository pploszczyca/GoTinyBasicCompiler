package utils

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
)

func PrintProgramTree(tree *domain.ProgramTree) {
	for _, node := range tree.Nodes {
		printNode(node, 0)
	}
}

func printNode(node *domain.Node, depth int) {
	prefix := getPrefix(depth)

	if len(node.Children) > 0 {
		fmt.Printf("%sNode Type: %d\n", prefix, node.Type)
		for _, child := range node.Children {
			printNode(child, depth+1)
		}
	} else {
		fmt.Printf("%sToken Type: %d, Value: %s\n", prefix, node.Token.Type, node.Token.Value)
	}
}

func getPrefix(depth int) string {
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "--"
	}
	return prefix
}
