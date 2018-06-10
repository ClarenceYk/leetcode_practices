package main

func levelOrderTraverse(root *TreeNode) [][]int {
	out := make([][]int, 0)
	for i := 1; i <= height(root); i++ {
		nodes := make([]int, 0)
		nodesOnGivenHight(root, i, &nodes)
		out = append(out, nodes)
	}
	return out
}

func height(root *TreeNode) int {
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

func nodesOnGivenHight(root *TreeNode, h int, nodes *[]int) {
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
