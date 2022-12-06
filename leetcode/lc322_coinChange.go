package leetcode

import (
	"fmt"
	"golang-acm/util"
	"math"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。

 */

func CoinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount+1
		fmt.Println()
		for j :=0; j < len(coins);  j++ {
			fmt.Print(j, ", ")
			//coins = [1, 2, 5], amount = 11
			//dp[i] = Min(dp[i], dp[i-coins[j]] + 1)
			//0, 0, 1, 0, 1, 0, 1, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 3
			//0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 0, 1, 2, 3
			if i >= coins[j] {
				dp[i] = util.Min(dp[i], dp[i-coins[j]] + 1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinDp(n int, coins []int) int {
	if n == 0{
		return 0
	}
	if n < 0 {
		return -1
	}
	//求最小值，所以初始化为正无穷
	res := math.MaxInt32
	for _, coin := range coins {
		subproblem := coinDp(n - coin, coins)
		//子问题无解，跳过
		if subproblem == -1{
			continue
		}
		res = util.Min(res, 1 + subproblem)
	}

	if res != math.MaxInt32 {
		return res
	} else {
		return -1
	}
}

func CoinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	n := len(coins)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = util.Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if amount < dp[amount] {
		return -1
	}

	return dp[amount]
}
