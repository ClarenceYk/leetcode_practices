package main

import "fmt"

// TreeNode example structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := getBinaryTree()
	fmt.Printf("Recursively preorder traverse binary tree: %v\n", preorderRecursiveTraverse(root))
	fmt.Printf("Iteratively preorder traverse binary tree: %v\n", preorderRecursiveTraverse(root))
}

// Binary tree example:
//        1
//       / \
//      2   3
//     / \ /
//    4  5 7
//        \
//         6
func getBinaryTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Left = &TreeNode{7, nil, nil}
	root.Left.Right.Right = &TreeNode{6, nil, nil}
	return root
}
