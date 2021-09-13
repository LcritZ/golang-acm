package leetcode

import (
	"golang-acm/util"
)

var inorderArr = []int{}

func InorderTraversal(root *util.TreeNode) []int {
	if root != nil {
		inorderArr = []int{}
		treeTraverse(root)
	}
	return inorderArr
}

func treeTraverse(root *util.TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		treeTraverse(root.Left)
	}
	inorderArr = append(inorderArr, root.Val)
	if root.Right != nil {
		treeTraverse(root.Right)
	}

	return
}

func InorderTraversal2Iteration(root *util.TreeNode) []int {
	if root != nil {
		inorderArr = []int{}
		treeTraverse(root)
	}
	return inorderArr
}

func GF_inorderTraversal(root *util.TreeNode) (res []int) {
	stack := []*util.TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
