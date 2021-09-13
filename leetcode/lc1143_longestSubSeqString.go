package leetcode

import (
	"fmt"
	"golang-acm/util"
)

/**
最长公共子序列 不要求连续，序列一样即可
abceg
xacg00

dp[1][0] = 0
dp[1][1] = 0
dp[1][2] = 0
dp[1][3] = 0
dp[1][4] = 0

dp[2][1] = 1
dp[2][2] = 1
dp[2][3] = 1
dp[2][4] = 1

dp[3][1] = 1
dp[3][2] = 1
dp[3][3] = 2  = dp[2][2] +1
dp[3][1] = 1




 */

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = util.Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

func LongestCommonSubSeq2(str1, str2 string) int {
	m, n := len(str1), len(str2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := [][]int{}

	for i := 0; i <= m; i++ {
		dp = append(dp,make([]int, n+1))
	}
	//fmt.Println(dp)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				// 如果字符相等，则在原基础dp[i][j]上加一
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				// 不等的话， 找到上边或者左边dp的最大值
				dp[i][j] = util.Max(dp[i - 1][j], dp[i][j - 1])
			}
			fmt.Println(i, dp[i])
		}
	}
	//for _, arr := range dp {
	//	fmt.Println(arr)
	//}
	return dp[m][n]
}