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

func NewCompleteBinaryTree(index int, elements ...int) *treeNode {
	if index < len(elements) {
		node := &treeNode{data: elements[index]}
		node.left = &treeNode{data: elements[2*index+1]}
		node.right = &treeNode{data: elements[2*index+2]}
		return node
	}
	return nil
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

func (root *treeNode) CountNodes() int {
	if root.left != nil && root.right != nil {
		return 1 + root.left.CountNodes() + root.right.CountNodes()
	} else if root.left == nil && root.right != nil {
		return 1 + root.right.CountNodes()
	} else if root.right == nil && root.left != nil {
		return 1 + root.left.CountNodes()
	}
	return 1
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

func (root *treeNode) IsComplete(index int, maxIndex int) bool {
	if index >= maxIndex {
		return false
	}
	if root.left != nil && root.right != nil {
		return root.left.IsComplete(2*index+1, maxIndex) && root.right.IsComplete(2*index+2, maxIndex)
	} else if root.right == nil && root.left != nil {
		return root.left.IsComplete(2*index+1, maxIndex)
	} else if root.left == nil && root.right != nil {
		return root.right.IsComplete(2*index+2, maxIndex)
	}
	return true
}

// Height balanced tree check
func (root *treeNode) IsBalanced(height int) bool {
	if root.left != nil && root.right != nil {
		return root.left.IsBalanced(height) && root.right.IsBalanced(height)
	} else if root.left != nil && root.right == nil {
		return root.left.IsBalanced(height + 1)
	} else if root.right != nil && root.left == nil {
		return root.right.IsBalanced(height + 1)
	} else {
		return height <= 1
	}
}
func (root *treeNode) IsBinarySearchTree() bool {
	left := root.left
	right := root.right
	if left != nil && right != nil {
		return left.data <= root.data && root.data <= right.data && left.IsBinarySearchTree() && right.IsBinarySearchTree()
	} else if left == nil && right != nil {
		return root.data <= right.data && right.IsBinarySearchTree()
	} else if right == nil && left != nil {
		return root.data <= left.data && left.IsBinarySearchTree()
	} else {
		return true
	}
}

func (root *treeNode) BinarySearch(el int) bool {
	if !root.IsBinarySearchTree() {
		fmt.Println("Not a binary search tree, expensive operation XD")
		return false
	}
	current := root.data
	if el == current {
		return true
	} else if current > el {
		if root.left != nil {
			return root.left.BinarySearch(el)
		}
		return false
	} else {
		if root.right != nil {
			return root.right.BinarySearch(el)
		}
		return false
	}
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
