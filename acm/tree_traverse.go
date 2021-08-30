package acm

import "golang-acm/common"

var treeArr []int

func PreOrder(root *common.TreeNode) []int {
	if root == nil {
		return treeArr
	}

	treeArr = append(treeArr, root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
	return treeArr
}

func InOrder(root *common.TreeNode) []int {
	if root == nil {
		return treeArr
	}

	InOrder(root.Left)
	treeArr = append(treeArr, root.Val)
	InOrder(root.Right)
	return treeArr
}

func PostOrder(root *common.TreeNode) []int {
	if root == nil {
		return treeArr
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	treeArr = append(treeArr, root.Val)
	return treeArr
}

