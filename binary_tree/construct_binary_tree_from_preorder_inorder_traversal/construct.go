package constructBinaryTreeFromPreorderInorderTraversal

import (
	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

// BuildTree construct a binary tree from preorder and inorder traversal
func BuildTree(preorder, inorder []int) *bt.TreeNode {
	if len(preorder) != len(inorder) || len(preorder) < 1 {
		return nil
	}

	root := &bt.TreeNode{
		Val: preorder[0],
	}

	pos := findValIn(inorder, preorder[0])

	inLeft := inorder[:pos]
	buildtreeHelper(&root.Left, preorder[1:1+len(inLeft)], inLeft)
	inRight := inorder[pos+1:]
	buildtreeHelper(&root.Right, preorder[len(preorder)-len(inRight):], inRight)

	return root
}

func buildtreeHelper(root **bt.TreeNode, pre, in []int) {
	if len(pre) == 0 || len(in) == 0 {
		return
	}

	*root = &bt.TreeNode{
		Val: pre[0],
	}

	pos := findValIn(in, pre[0])

	inLeft := in[:pos]
	buildtreeHelper(&((*root).Left), pre[1:1+len(inLeft)], inLeft)
	inRight := in[pos+1:]
	buildtreeHelper(&((*root).Right), pre[len(pre)-len(inRight):], inRight)
}

func findValIn(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}
