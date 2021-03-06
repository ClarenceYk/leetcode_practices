package traverse

import bt "github.com/clarenceyk/leetcode_practices/binary_tree"

// LevelOrderTraverse traverse a tree in level order
func LevelOrderTraverse(root *bt.TreeNode) [][]int {
	out := make([][]int, 0)
	for i := 1; i <= height(root); i++ {
		nodes := make([]int, 0)
		nodesOnGivenHight(root, i, &nodes)
		out = append(out, nodes)
	}
	return out
}

func height(root *bt.TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	lr := height(root.Right)
	if lh > lr {
		return lh + 1
	}
	return lr + 1
}

func nodesOnGivenHight(root *bt.TreeNode, h int, nodes *[]int) {
	if root == nil {
		return
	}
	if h == 1 {
		*nodes = append(*nodes, root.Val)
	} else {
		nodesOnGivenHight(root.Left, h-1, nodes)
		nodesOnGivenHight(root.Right, h-1, nodes)
	}
}
