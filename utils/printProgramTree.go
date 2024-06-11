package utils

import (
	"GoTinyBasicCompiler/domain"
	"fmt"
)

func PrintProgramTree(tree *domain.ProgramTree) {
	for _, node := range tree.Nodes {
		printNode(node, "", true)
	}
}

func printNode(node *domain.Node, prefix string, isTail bool) {
	if len(node.Children) > 0 {
		fmt.Printf("%s%sNODE: %s\n", prefix, getBranch(isTail), node.Type)
		newPrefix := getNewPrefix(prefix, isTail)
		for i, child := range node.Children {
			printNode(child, newPrefix, i == len(node.Children)-1)
		}
	} else {
		fmt.Printf("%s%sTOKEN Type: %s, Value: %s\n", prefix, getBranch(isTail), node.Token.Type, node.Token.Value)
	}
}

func getBranch(isTail bool) string {
	if isTail {
		return "└── "
	}
	return "├── "
}

func getNewPrefix(prefix string, isTail bool) string {
	if isTail {
		return prefix + "    "
	}
	return prefix + "│   "
}
