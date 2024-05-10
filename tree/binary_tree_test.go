package tree

import (
	"testing"
)

func createSampleTree() *treeNode {
	/*
		This is the right side of the tree

			│   │   ├── 7
		│   ├── 2
		└── 1
			└── 3
				└── 4

		This is the left side of the tree
	*/
	tree := treeNode{data: 1}
	tree.right = &treeNode{data: 2}
	tree.left = &treeNode{data: 3}
	tree.right.right = &treeNode{data: 7}
	tree.left.left = &treeNode{data: 4}
	return &tree
}
func createFullTree() *treeNode {
	tree := treeNode{data: 1}
	tree.right = &treeNode{data: 2}
	tree.left = &treeNode{data: 3}
	tree.right.right = &treeNode{data: 7}
	tree.right.left = &treeNode{data: 8}
	return &tree
}

func createPerfectTree() *treeNode {
	tree := treeNode{data: 1}
	tree.right = &treeNode{data: 2}
	tree.left = &treeNode{data: 3}
	tree.right.right = &treeNode{data: 7}
	tree.right.left = &treeNode{data: 8}
	tree.left.left = &treeNode{data: 4}
	tree.left.right = &treeNode{data: 5}
	return &tree
}

func TestInOrderTraversal(t *testing.T) {
	tree := createSampleTree()
	elements := tree.InOrderTraversal()
	expected := []int{4, 3, 1, 2, 7}
	for i, el := range elements {
		if expected[i] != el {
			t.Fatalf("inorder traversal wrong, expected : %d, but got : %d\n", expected[i], el)
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := createSampleTree()
	elements := tree.PreOrderTraversal()
	expected := []int{1, 3, 4, 2, 7}
	for i, el := range elements {
		if expected[i] != el {
			t.Fatalf("preorders traversal wrong, expected : %d, but got : %d\n", expected[i], el)
		}
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := createSampleTree()
	elements := tree.PostOrderTraversal()
	expected := []int{4, 3, 7, 2, 1}
	for i, el := range elements {
		if expected[i] != el {
			t.Fatalf("postorder traversal wrong, expected : %d, but got : %d\n", expected[i], el)
		}
	}
}

func TestTreeHeight(t *testing.T) {
	tree := createSampleTree()
	if tree.Height() != 2 {
		t.Fatal("wrong height")
	}
}

func TestTreeIsFull(t *testing.T) {
	tree := createSampleTree()
	fullTree := createFullTree()
	if tree.IsFull() {
		t.Fatalf("wrong the tree is not full \n%s", tree)
	}
	if !fullTree.IsFull() {
		t.Fatalf("wrong the tree is full \n%s", fullTree)

	}
}

func TestTreeIsPerfect(t *testing.T) {
	tree := createSampleTree()
	perfectTree := createPerfectTree()
	if tree.IsPerfect(tree.Height(), 0) {
		t.Fatalf("wrong the tree is not perfect \n%s", tree)
	}
	if !perfectTree.IsPerfect(perfectTree.Height(), 0) {
		t.Fatalf("wrong the tree is perfect \n%s", perfectTree)
	}
}
