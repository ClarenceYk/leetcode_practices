package main

import (
	"fmt"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
	"github.com/clarenceyk/leetcode_practices/binary_tree/traverse_a_tree"
)

func main() {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 3, 7, 20, 15})
	fmt.Println(traverse.LevelOrderTraverse(root))
}

func buildTree(inorder, postorder []int) *bt.TreeNode {
	if (len(inorder) != len(postorder)) || len(inorder) < 1 {
		return nil
	}

	tail := len(postorder) - 1
	rootVal := postorder[tail]
	postorder = postorder[:tail]

	root := &bt.TreeNode{
		Val: rootVal,
	}

	top := findValIn(inorder, rootVal)
	buildTreeHelper(&root.Left, inorder[:top], postorder[:top])
	inRight := inorder[top+1:]
	buildTreeHelper(&root.Right, inRight, postorder[len(postorder)-len(inRight):])
	return root
}

func buildTreeHelper(root **bt.TreeNode, in, post []int) {
	if (len(in) != len(post)) || len(in) < 1 {
		return
	}

	tail := len(post) - 1
	rootVal := post[tail]
	post = post[:tail]

	*root = &bt.TreeNode{
		Val: rootVal,
	}

	top := findValIn(in, rootVal)
	buildTreeHelper(&((*root).Left), in[:top], post[:top])
	inRight := in[top+1:]
	buildTreeHelper(&((*root).Right), inRight, post[len(post)-len(inRight):])
}

func findValIn(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return -1
}
