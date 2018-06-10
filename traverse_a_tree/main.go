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
	fmt.Println()

	fmt.Printf("Recursively inorder traverse binary tree: %v\n", inorderRecursiveTraverse(root))
	fmt.Printf("Iteratively inorder traverse binary tree: %v\n", inorderIterativeTraverse(root))
	fmt.Println()

	fmt.Printf("Recursively postorder traverse binary tree:    %v\n", postorderRecursiveTraverse(root))
	fmt.Printf("Iteratively postorder traverse binary tree(1): %v\n", postorderIterativeTraverse1(root))
	fmt.Printf("Iteratively postorder traverse binary tree(2): %v\n", postorderIterativeTraverse2(root))
	fmt.Println()

	fmt.Printf("Level order traverse binary tree: %v\n", levelOrderTraverse(root))
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
