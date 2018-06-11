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
	fmt.Printf("Is this binary tree is a mirror of itself(recursive version)? %v\n", isSymmetricTreeRecursiveVersion(root))
	fmt.Printf("Is this binary tree is a mirror of itself(iterative version)? %v\n", isSymmetricTreeIterativeVersion(root))
}

// Tree example
//       1
//      / \
//     2   2
//    / \ / \
//   3  4 4  3
func getTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	root.Right.Left = &TreeNode{4, nil, nil}
	root.Right.Right = &TreeNode{3, nil, nil}
	return root
}

func isSymmetricTreeRecursiveVersion(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTreeRecursiveHelper(root.Left, root.Right)
}

func isSymmetricTreeRecursiveHelper(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricTreeRecursiveHelper(left.Left, right.Right) &&
		isSymmetricTreeRecursiveHelper(left.Right, right.Left)
}

func isSymmetricTreeIterativeVersion(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := make([]*TreeNode, 1)
	top := 0

	if root.Left != nil {
		if root.Right == nil {
			return false
		}
		q = append(q, root.Left)
		q = append(q, root.Right)
		top += 2
	} else if root.Right != nil {
		return false
	}

	for top > 0 {
		right := q[top]
		q = q[:top]
		top--

		left := q[top]
		q = q[:top]
		top--

		if left.Val != right.Val {
			return false
		}

		if left.Left != nil {
			if right.Right == nil {
				return false
			}
			q = append(q, left.Left)
			q = append(q, right.Right)
			top += 2
		} else if right.Right != nil {
			return false
		}

		if left.Right != nil {
			if right.Left == nil {
				return false
			}
			q = append(q, left.Right)
			q = append(q, right.Left)
			top += 2
		} else if right.Left != nil {
			return false
		}
	}

	return true
}
