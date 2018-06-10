package main

func inorderRecursiveTraverse(root *TreeNode) []int {
	out := make([]int, 0)
	inorderRecursiveTraverseHelper(root, &out)
	return out
}

func inorderIterativeTraverse(root *TreeNode) []int {
	return inorderIterativeTraverseHelper(root)
}

func inorderRecursiveTraverseHelper(root *TreeNode, out *[]int) {
	if root == nil {
		return
	}
	inorderRecursiveTraverseHelper(root.Left, out)
	*out = append(*out, root.Val)
	inorderRecursiveTraverseHelper(root.Right, out)
}

func inorderIterativeTraverseHelper(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*TreeNode, 1)
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
