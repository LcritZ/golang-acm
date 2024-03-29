package basic

import (
	"golang-acm/util"
	"math"
)

/**

二叉树最长路径

 */

var longestPath int

func LongestPath(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	longestPath = math.MinInt64
	LongestPathHelper(root)
	return longestPath
}

func LongestPathHelper(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	leftPath := util.Max(LongestPathHelper(root.Left), 0)
	rightPath := util.Max(LongestPathHelper(root.Right), 0)
	longestPath = util.Max(leftPath+rightPath+1, longestPath)
	curSum := util.Max(leftPath, leftPath)
	curSum += 1
	return curSum
}
