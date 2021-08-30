package acm

import (
	"golang-acm/common"
	"golang-acm/leetcode"
	"math"
)

/**

二叉树最长路径

 */

var longestPath int

func LongestPath(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	longestPath = math.MinInt64
	LongestPathHelper(root)
	return longestPath
}

func LongestPathHelper(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftPath := common.Max(LongestPathHelper(root.Left), 0)
	rightPath := common.Max(LongestPathHelper(root.Right), 0)
	leetcode.MaxSum = common.Max(leftPath+rightPath+1, leetcode.MaxSum)
	curSum := common.Max(leftPath, leftPath)
	curSum += 1
	return curSum
}
