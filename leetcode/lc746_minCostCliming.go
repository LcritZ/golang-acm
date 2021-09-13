package leetcode

import (
	"golang-acm/util"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, util.Min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = util.Min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
