package tree

import (
	"fmt"
	"math"
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
	elements := make([]int, 0)

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
	elements := make([]int, 0)

	elements = append(elements, root.data)

	if root.left != nil {
		elements = append(elements, root.left.PreOrderTraversal()...)
	}

	if root.right != nil {
		elements = append(elements, root.right.PreOrderTraversal()...)
	}

	return elements
}

func (root *treeNode) PostOrderTraversal() []int {
	elements := make([]int, 0)

	if root.left != nil {
		elements = append(elements, root.left.PostOrderTraversal()...)
	}
	if root.right != nil {
		elements = append(elements, root.right.PostOrderTraversal()...)
	}

	elements = append(elements, root.data)

	return elements

}
func (root *treeNode) Height() int {
	if root.left == nil && root.right == nil {
		return 0
	} else if root.left == nil {
		return root.right.Height() + 1
	} else if root.right == nil {
		return root.left.Height() + 1
	}
	height := math.Max(float64(root.left.Height()), float64(root.right.Height())) + 1
	return int(height)
}

func (root *treeNode) IsFull() bool {
	// internal nodes either have 2 child nodes or none
	if root.left == nil && root.right == nil {
		return true
	} else if root.left != nil && root.right != nil {
		return root.left.IsFull() && root.right.IsFull()
	}
	return false
}

func (root *treeNode) IsPerfect(height int, currentLvL int) bool {
	// internal nodes have exactly 2 child nodes
	if root.left == nil && root.right == nil {
		return currentLvL == height
	} else if root.left == nil || root.right == nil {
		return false
	}

	return root.left.IsPerfect(height, currentLvL+1) && root.right.IsPerfect(height, currentLvL+1)
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
