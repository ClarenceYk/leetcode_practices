package constructBinaryTreeFromPreorderInorderTraversal

import (
	"testing"

	"github.com/clarenceyk/leetcode_practices/binary_tree/traverse_a_tree"

	bt "github.com/clarenceyk/leetcode_practices/binary_tree"
)

type input struct {
	pre []int
	in  []int
}

func TestBuildTree(t *testing.T) {
	tables := []struct {
		inp  input
		outp *bt.TreeNode
	}{
		{
			inp:  input{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			outp: getResultTree(),
		},
	}

	for _, table := range tables {
		tree := BuildTree(table.inp.pre, table.inp.in)
		if !treeEqual(tree, table.outp) {
			t.Errorf("expect:%v get:%v\n",
				traverse.LevelOrderTraverse(table.outp),
				traverse.LevelOrderTraverse(tree))
		}
	}
}

func treeEqual(t1, t2 *bt.TreeNode) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}
	return t1.Val == t2.Val && treeEqual(t1.Left, t2.Left) && treeEqual(t1.Right, t2.Right)
}

func getResultTree() *bt.TreeNode {
	root := &bt.TreeNode{
		Val: 3,
	}
	root.Left = &bt.TreeNode{
		Val: 9,
	}
	root.Right = &bt.TreeNode{
		Val: 20,
	}
	root.Right.Left = &bt.TreeNode{
		Val: 15,
	}
	root.Right.Right = &bt.TreeNode{
		Val: 7,
	}
	return root
}
