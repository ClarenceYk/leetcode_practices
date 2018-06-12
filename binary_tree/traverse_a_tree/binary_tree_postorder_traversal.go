package traverse

import bt "github.com/clarenceyk/leetcode_practices/binary_tree"

func postorderRecursiveTraverse(root *bt.TreeNode) []int {
	out := make([]int, 0)
	postorderRecursiveTraverseHelper(root, &out)
	return out
}

func postorderIterativeTraverse1(root *bt.TreeNode) []int {
	return postorderIterativeTraverseHelper1(root)
}

func postorderIterativeTraverse2(root *bt.TreeNode) []int {
	return postorderIterativeTraverseHelper2(root)
}

func postorderRecursiveTraverseHelper(root *bt.TreeNode, out *[]int) {
	if root == nil {
		return
	}
	postorderRecursiveTraverseHelper(root.Left, out)
	postorderRecursiveTraverseHelper(root.Right, out)
	*out = append(*out, root.Val)
}

func postorderIterativeTraverseHelper1(root *bt.TreeNode) []int {
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
		if root.Left != nil {
			q = append(q, root.Left)
			top++
		}
		if root.Right != nil {
			q = append(q, root.Right)
			top++
		}
	}

	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-i-1] = out[len(out)-i-1], out[i]
	}

	return out
}

func postorderIterativeTraverseHelper2(root *bt.TreeNode) []int {
	if root == nil {
		return nil
	}

	q := make([]*bt.TreeNode, 1)
	top := 0

	if root.Right != nil {
		q = append(q, root.Right)
		top++
	}
	q = append(q, root)
	top++
	root = root.Left

	out := make([]int, 0)
	for top > 0 {
		for root != nil {
			if root.Right != nil {
				q = append(q, root.Right)
				top++
			}
			q = append(q, root)
			top++
			root = root.Left
		}
		root = q[top]
		q = q[:top]
		top--
		if root.Right != nil && top > 0 && q[top] == root.Right {
			tmp := q[top]
			q = q[:top]
			// top--
			q = append(q, root)
			// top++
			root = tmp
		} else {
			out = append(out, root.Val)
			root = nil
		}
	}

	return out
}
