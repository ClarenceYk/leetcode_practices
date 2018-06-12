package main

import (
	"fmt"
)

// TreeNode example
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := getTree()
	fmt.Printf("Dose this binary tree has a path sum? %v\n", hasPathSum(root, 22))
}

// Tree example
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
func getTree() *TreeNode {
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{8, nil, nil}
	root.Left.Left = &TreeNode{11, nil, nil}
	root.Right.Left = &TreeNode{13, nil, nil}
	root.Right.Right = &TreeNode{4, nil, nil}
	root.Left.Left.Left = &TreeNode{7, nil, nil}
	root.Left.Left.Right = &TreeNode{2, nil, nil}
	root.Right.Right.Right = &TreeNode{1, nil, nil}
	return root
}

func hasPathSum(root *TreeNode, sum int) bool {
	cmp := 0
	return hasPathSumHelper(root, sum, &cmp)
}

func hasPathSumHelper(root *TreeNode, sum int, cmp *int) bool {
	if root == nil {
		return false
	}
	*cmp += root.Val
	if hasPathSumHelper(root.Left, sum, cmp) || hasPathSumHelper(root.Right, sum, cmp) {
		return true
	}
	if root.Left == nil && root.Right == nil && sum == *cmp {
		return true
	}
	*cmp -= root.Val
	return false
}
