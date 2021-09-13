package basic

import (
	"fmt"
	"golang-acm/util"
)

/**
背包问题


 */

func BackPack(nums [][]int, total int) int {
	dp := make([][]int, len(nums))
	//初始化二维数组
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, total+1)
	}

	//放入第一个物品，填第一行列表
	for i:= nums[0][0]; i < total; i++ {
		dp[0][i] = nums[0][1]
	}

	for i := 1; i < len(nums); i++ {
		for j:= nums[i][0]; j < total; j++ {
			dp[i][j] = util.Max(dp[i-1][j], dp[i-1][j-nums[i][0]] + nums[i][1])
		}
	}
	return dp[len(nums) - 1][total]
}

func ZeroOnePack(weight []int, value []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i := weight[0]; i < capacity; i++ {
		dp[i] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := capacity; j >= 0; j-- {
			if j >= weight[i] {
				fmt.Println(j, dp[j], dp[j-weight[i]], value[i])
				dp[j] = util.Max(dp[j], dp[j-weight[i]]+value[i])
				fmt.Println(j, dp[j], "--")
			}
		}
	}

	return dp[capacity]
}

func ZeroOnePack2(weight []int, value []int, cap int) int {
	dp := make([]int, cap +1)
	//for i := 0; i < cap; i++ {
	//	dp[i] = 0
	//}
	/**
	size := []int{2, 3, 4, 5, 7} //物品大小
	value := []int{2, 4, 4, 7, 9} //物品价值
	 */

	for i := 0; i < len(weight); i++ {
		fmt.Println(weight[i], "--")
		for j := cap; j >= weight[i]; j-- {
			fmt.Println(j, dp[j], dp[j-weight[i]], value[i])
			dp[j] = util.Max(dp[j], dp[j-weight[i]] + value[i])
			fmt.Println(dp[j], "-")
		}
	}

	return dp[cap]
}

func TotalPack(weight []int, value []int, cap int) int {
	dp := make([]int, cap +1)

	/**
	size := []int{2, 3, 4, 5, 7} //物品大小
	value := []int{2, 4, 4, 7, 9} //物品价值
	*/

	for i := 0; i < len(weight); i++ {
		fmt.Println(weight[i], "--")
		for j := weight[i]; j <= cap; j++ {
			fmt.Println(j, dp[j], dp[j-weight[i]], value[i])
			dp[j] = util.Max(dp[j], dp[j-weight[i]]+value[i])
			fmt.Println(dp[j], "-")
		}
	}

	return dp[cap]

}