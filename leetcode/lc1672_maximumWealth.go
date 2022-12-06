package leetcode

import "golang-acm/util"

func maximumWealth(accounts [][]int) int {
    max := 0
    for i := 0; i < len(accounts); i++ {
        curr := 0
        for j := 0; j < len(accounts[0]); j++ {
            curr += accounts[i][j]
        }
        max = util.Max(max, curr)
    }

    return max
}
