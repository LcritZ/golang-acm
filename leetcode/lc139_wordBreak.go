package leetcode

import fmt "fmt"

func WordBreak(s string, wordDict []string) bool {
    if len(s) == 0 {
        return true
    }

    wordMap := make(map[string]bool)
    for i := 0; i < len(wordDict); i++ {
        wordMap[wordDict[i]] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true

    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordMap[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }

    fmt.Println(dp)
    return dp[len(s)]

}

