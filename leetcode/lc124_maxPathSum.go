package leetcode

import (
	"golang-acm/common"
	"math"
)

var MaxSum int


/**
最大路径和
后序遍历思想


 */
func MaxPathSum(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	MaxSum = math.MinInt64
	MaxPathSumHelper(root)
	return MaxSum
}

func MaxPathSumHelper(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := common.Max(MaxPathSumHelper(root.Left), 0)
	rightSum := common.Max(MaxPathSumHelper(root.Right), 0)
	MaxSum = common.Max(leftSum+rightSum+root.Val, MaxSum)
	curSum := common.Max(leftSum, rightSum)
	curSum += root.Val
	return curSum
}
