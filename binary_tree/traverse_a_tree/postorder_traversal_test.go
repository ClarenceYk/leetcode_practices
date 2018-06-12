package traverse

import (
	"testing"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

// TestPostorderRecursiveTraverse test traversing a tree recursively in postorder
func TestPostorderRecursiveTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{4, 6, 5, 2, 7, 3, 1}},
	}

	for _, table := range tables {
		arr := postorderRecursiveTraverse(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}

// TestPostorderIterativeTraverse1 test traversing a tree iteratively in postorder using method 1
func TestPostorderIterativeTraverse1(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{4, 6, 5, 2, 7, 3, 1}},
	}

	for _, table := range tables {
		arr := postorderIterativeTraverse1(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}

// TestPostorderIterativeTraverse2 test traversing a tree iteratively in postorder using method 2
func TestPostorderIterativeTraverse2(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out []int
	}{
		{bt.GetBinaryTree(), []int{4, 6, 5, 2, 7, 3, 1}},
	}

	for _, table := range tables {
		arr := postorderIterativeTraverse2(table.in)
		for i, v := range arr {
			if v != table.out[i] {
				t.Errorf("Want %v; Get %v\n", table.out, arr)
				break
			}
		}
	}
}
