package traverse

import bt "github.com/clarenceyk/leetcode_practices/binary_tree"

func inorderRecursiveTraverse(root *bt.TreeNode) []int {
	out := make([]int, 0)
	inorderRecursiveTraverseHelper(root, &out)
	return out
}

func inorderIterativeTraverse(root *bt.TreeNode) []int {
	return inorderIterativeTraverseHelper(root)
}

func inorderRecursiveTraverseHelper(root *bt.TreeNode, out *[]int) {
	if root == nil {
		return
	}
	inorderRecursiveTraverseHelper(root.Left, out)
	*out = append(*out, root.Val)
	inorderRecursiveTraverseHelper(root.Right, out)
}

func inorderIterativeTraverseHelper(root *bt.TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*bt.TreeNode, 1)
	top := 0

	for root != nil {
		q = append(q, root)
		top++
		root = root.Left
	}

	out := make([]int, 0)
	for top > 0 {
		root = q[top]
		q = q[:top]
		top--

		out = append(out, root.Val)
		if root = root.Right; root != nil {
			for root != nil {
				q = append(q, root)
				top++
				root = root.Left
			}
		}
	}

	return out
}
