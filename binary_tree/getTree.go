package binaryTree

// GetBinaryTree get a binary tree like this:
//         1
//       /  \
//      2    3
//     / \  /
//    4  5  7
//       \
//       6
func GetBinaryTree() *TreeNode {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left.Left = &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right = &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	root.Left.Right.Right = &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	return root
}
