package traverse

import (
	"testing"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

// TestPreorderRecursiveTraverse test traversing a tree recursively in preorder
func TestPreorderRecursiveTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{1, 2, 4, 5, 6, 3, 7}},
	}

	for _, table := range tables {
		arr := preorderRecursiveTraverse(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}

// TestPreorderIterativeTraverse test traversing a tree iteratively in preorder
func TestPreorderIterativeTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{1, 2, 4, 5, 6, 3, 7}},
	}

	for _, table := range tables {
		arr := preorderIterativeTraverse(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}
