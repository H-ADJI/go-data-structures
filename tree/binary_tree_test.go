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

func createCompleteTree() *treeNode {
	tree := treeNode{data: 1}
	tree.right = &treeNode{data: 2}
	tree.left = &treeNode{data: 3}
	tree.right.left = &treeNode{data: 8}
	tree.left.left = &treeNode{data: 4}
	tree.left.right = &treeNode{data: 5}
	return &tree
}

func createBalancedTree() *treeNode {
	tree := treeNode{data: 1}
	tree.right = &treeNode{data: 2}
	tree.left = &treeNode{data: 3}
	tree.right.left = &treeNode{data: 8}
	tree.left.left = &treeNode{data: 4}
	tree.left.right = &treeNode{data: 5}
	return &tree
}

func createUnBalancedTree() *treeNode {
	tree := treeNode{data: 1}
	tree.left = &treeNode{data: 3}
	tree.left.left = &treeNode{data: 4}
	tree.left.right = &treeNode{data: 5}
	return &tree
}

func createBinarySearchTree() *treeNode {
	tree := treeNode{data: 8}
	tree.right = &treeNode{data: 10}
	tree.left = &treeNode{data: 5}
	tree.right.left = &treeNode{data: 9}
	tree.right.right = &treeNode{data: 11}
	tree.left.left = &treeNode{data: 4}
	tree.left.right = &treeNode{data: 7}
	tree.left.right.left = &treeNode{data: 6}
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

func TestTreeIsComplete(t *testing.T) {
	tree := createSampleTree()
	completeTree := createCompleteTree()
	if tree.IsComplete(0, tree.CountNodes()) {
		t.Fatalf("wrong the tree is not complete \n%s", tree)
	}
	if !completeTree.IsComplete(0, completeTree.CountNodes()) {
		t.Fatalf("wrong the tree is complete \n%s", completeTree)
	}
}

func TestTreeNodesCount(t *testing.T) {
	tree := createSampleTree()
	if tree.CountNodes() != 5 {
		t.Fatalf("wrong nubmer of nodes in the tree %d \n", tree.CountNodes())
	}
}

func TestTreeIsBalanced(t *testing.T) {
	unBalancedTree := createUnBalancedTree()
	balancedTree := createBalancedTree()
	isbalanced, _ := unBalancedTree.IsBalanced()
	if isbalanced {
		t.Fatalf("wrong the tree is not balanced \n%s", unBalancedTree)
	}
	isbalanced, _ = balancedTree.IsBalanced()
	if !isbalanced {
		t.Fatalf("wrong the tree is balanced \n%s", balancedTree)
	}
}

func TestIsBinarySearchTree(t *testing.T) {
	tree := createSampleTree()
	binarySearchTree := createBinarySearchTree()
	if tree.IsBinarySearchTree() {
		t.Fatalf("wrong the tree is not a binary search tree \n%s", tree)
	}
	if !binarySearchTree.IsBinarySearchTree() {
		t.Fatalf("wrong the tree is a bianry search tree \n%s", binarySearchTree)
	}
}

func TestBinarySearch(t *testing.T) {
	binarySearchTree := createBinarySearchTree()
	tree := createSampleTree()
	if !binarySearchTree.BinarySearch(10) {
		t.Fatalf("wrong the tree contains the node 10 \n%s", binarySearchTree)
	}
	if binarySearchTree.BinarySearch(100) {
		t.Fatalf("wrong the tree doesn't contains the node 100 \n%s", binarySearchTree)
	}
	if tree.BinarySearch(100) {
		t.Fatalf("wrong, impossible operation the tree isn't a bianry search tree \n%s", binarySearchTree)
	}
}

func TestInsertIntoBinarySearchTree(t *testing.T) {
	binarySearchTree := createBinarySearchTree()
	binarySearchTree.InsertIntoBinarySearchTree(99)
	if !binarySearchTree.IsBinarySearchTree() {
		t.Fatalf("wrong the tree is not a binary search tree \n%s", binarySearchTree)
	}
	if !binarySearchTree.BinarySearch(99) {
		t.Fatalf("wrong the tree contains the node 99 \n%s", binarySearchTree)
	}

	if !binarySearchTree.IsBinarySearchTree() {
		t.Fatalf("wrong the tree is not a binary search tree \n%s", binarySearchTree)
	}
	if !binarySearchTree.BinarySearch(10) {
		t.Fatalf("wrong the tree contains the node 10 \n%s", binarySearchTree)
	}
}

func TestDeleteFromBST(t *testing.T) {
	binarySearchTree := createBinarySearchTree()
	toFind := []int{5, 10, 4, 11, 9}
	notFind := []int{8, 7, 6, 7, 6, 7}
	for _, el := range notFind {
		binarySearchTree = binarySearchTree.DeleteFromBinarySearchTree(el)
		if binarySearchTree.BinarySearch(el) {
			t.Fatalf("wrong, the tree shouldn't contain %d \n%s", el, binarySearchTree)
		}
	}
	for _, el := range toFind {
		if !binarySearchTree.BinarySearch(el) {
			t.Fatalf("wrong, the tree should contain %d \n%s", el, binarySearchTree)
		}
	}

}
