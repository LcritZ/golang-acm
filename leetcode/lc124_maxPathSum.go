package leetcode

import (
	"golang-acm/util"
	"math"
)

var MaxSum int


/**
最大路径和
后序遍历思想


 */
func MaxPathSum(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	MaxSum = math.MinInt64
	MaxPathSumHelper(root)
	return MaxSum
}

func MaxPathSumHelper(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := util.Max(MaxPathSumHelper(root.Left), 0)
	rightSum := util.Max(MaxPathSumHelper(root.Right), 0)
	MaxSum = util.Max(leftSum+rightSum+root.Val, MaxSum)
	curSum := util.Max(leftSum, rightSum)
	curSum += root.Val
	return curSum
}
