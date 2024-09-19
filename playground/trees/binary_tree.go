package playground

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func printIndent(indent int) {
	fmt.Print(strings.Repeat(" ", indent))
}

func TreeDisplay(node *TreeNode, indent int) {
	if node != nil {
		if node.right != nil {
			TreeDisplay(node.right, indent+4)
		}

		if indent > 0 {
			printIndent(indent)
		}

		if node.right != nil {
			fmt.Println(" /")
			printIndent(indent)
		}

		fmt.Println(node.data)

		if node.left != nil {
			printIndent(indent)
			fmt.Println(" \\")
			TreeDisplay(node.left, indent+4)
		}
	}
}
