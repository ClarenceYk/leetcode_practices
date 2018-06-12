package traverse

import (
	"testing"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

// TestLevelOrderTraverse test traversing a binary tree recursively in level order
func TestLevelOrderTraverse(t *testing.T) {
	tables := []struct {
		in  *bt.TreeNode
		out [][]int
	}{
		{bt.GetBinaryTree(), [][]int{
			[]int{1},
			[]int{2, 3},
			[]int{4, 5, 7},
			[]int{6},
		}},
	}

	for _, table := range tables {
		arr := levelOrderTraverse(table.in)
		for i, a := range arr {
			for j, v := range a {
				if v != table.out[i][j] {
					t.Errorf("Want %v; Get %v\n", table.out, arr)
					break
				}
			}
		}
	}
}
