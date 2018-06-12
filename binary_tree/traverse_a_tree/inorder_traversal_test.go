package traverse

import (
	"testing"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

// TestInorderRecursiveTraverse test traversing a tree recursively in inorder
func TestInorderRecursiveTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{4, 2, 5, 6, 1, 7, 3}},
	}

	for _, table := range tables {
		arr := inorderRecursiveTraverse(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}

// TestInorderIterativeTraverse test traversing a tree iteratively in inorder
func TestInorderIterativeTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{4, 2, 5, 6, 1, 7, 3}},
	}

	for _, table := range tables {
		arr := inorderIterativeTraverse(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}
