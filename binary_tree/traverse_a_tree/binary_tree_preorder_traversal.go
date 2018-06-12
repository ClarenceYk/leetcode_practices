package traverse

import bt "github.com/clarenceyk/leetcode_practices/binary_tree"

func preorderRecursiveTraverse(root *bt.TreeNode) []int {
	out := make([]int, 0)
	preorderRecursiveTraverseHelper(root, &out)
	return out
}

func preorderIterativeTraverse(root *bt.TreeNode) []int {
	return preorderIterativeTraverseHelper(root)
}

func preorderRecursiveTraverseHelper(root *bt.TreeNode, out *[]int) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)
	preorderRecursiveTraverseHelper(root.Left, out)
	preorderRecursiveTraverseHelper(root.Right, out)
}

func preorderIterativeTraverseHelper(root *bt.TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*bt.TreeNode, 1)
	top := 0

	q = append(q, root)
	top++

	out := make([]int, 0)
	for top > 0 {
		root = q[top]
		q = q[:top]
		top--

		out = append(out, root.Val)
		if root.Right != nil {
			q = append(q, root.Right)
			top++
		}
		if root.Left != nil {
			q = append(q, root.Left)
			top++
		}
	}

	return out
}
