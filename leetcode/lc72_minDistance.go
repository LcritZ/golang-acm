package leetcode

import "golang-acm/util"

func MinDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    dp[0][0] = 0
    for i := 1; i < m+1; i++ {
        for j := 1; j < n+1; j++ {
            min := util.Min(dp[i-1][j], dp[i][j-1])
            min = util.Min(min, dp[i-1][j-1])
            dp[i][j] = min+1
            if word1[i-1] == word2[j-1] {
                dp[i][j] = util.Min(dp[i][j], dp[i-1][j-1])
            }
        }
    }

    return dp[m][n]
}

