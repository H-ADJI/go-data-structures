package tree

import (
	"fmt"
	"math/rand"
	"strings"
)

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func NewRandomBinaryTree(elements ...int) *treeNode {
	var rootNode treeNode
	nbElements := len(elements)
	if nbElements == 0 {
		return nil
	}
	idx := rand.Intn(nbElements)
	rootNode.data = elements[idx]
	rootNode.left = NewRandomBinaryTree(elements[:idx]...)
	rootNode.right = NewRandomBinaryTree(elements[idx+1:]...)
	return &rootNode
}

func (root *treeNode) InOrderTraversal() []int {
	elements := make([]int, 1)

	if root.left != nil {
		elements = append(elements, root.left.InOrderTraversal()...)
	}

	elements = append(elements, root.data)

	if root.right != nil {
		elements = append(elements, root.right.InOrderTraversal()...)
	}

	return elements
}

func (root *treeNode) PreOrderTraversal() []int {
	elements := make([]int, 1)

	elements = append(elements, root.data)

	if root.left != nil {
		elements = append(elements, root.left.InOrderTraversal()...)
	}

	if root.right != nil {
		elements = append(elements, root.right.InOrderTraversal()...)
	}

	return elements
}

func (root *treeNode) PostOrderTraversal() []int {
	elements := make([]int, 1)

	if root.left != nil {
		elements = append(elements, root.left.InOrderTraversal()...)
	}
	if root.right != nil {
		elements = append(elements, root.right.InOrderTraversal()...)
	}

	elements = append(elements, root.data)

	return elements

}

func (root *treeNode) String() string {
	return root.stringWithIndent("", true)
}

func (node *treeNode) stringWithIndent(indent string, isTail bool) string {
	if node == nil {
		return ""
	}

	var builder strings.Builder

	if node.right != nil {
		builder.WriteString(node.right.stringWithIndent(indent+"│   ", false))
	}

	builder.WriteString(indent)
	if isTail {
		builder.WriteString("└── ")
	} else {
		builder.WriteString("├── ")
	}
	builder.WriteString(fmt.Sprintf("%d\n", node.data))

	if node.left != nil {
		builder.WriteString(node.left.stringWithIndent(indent+"    ", true))
	}

	return builder.String()
}
