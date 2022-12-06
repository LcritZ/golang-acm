package leetcode

import "fmt"

func PossibleBipartition(n int, dislikes [][]int) bool {
    if n == 1 {
        return false
    }

    visited := make([]int, n+1)
    groups := make([][]int, 2)
    groups[0] = []int{}
    groups[1] = []int{}
    //groups[3] = []int{}

    dislikeMap := make(map[int][]int)

    for i := 0; i < len(dislikes); i++ {
        dislikeMap[dislikes[i][0]] = append(dislikeMap[dislikes[i][0]], dislikes[i][1])
    }


    //sort.Slice(dislikes, func(i, j int) bool {
    //    if dislikes[i][0] < dislikes[j][0] {
    //        return true
    //    } else {
    //        return false
    //    }
    //})

    var dfs func(n int, x int)
    dfs = func(n int, x int) {
        if visited[n] != 0 {
            return
        }
        groups[x-1] = append(groups[x-1], n)
        visited[n] = x
        next := 0
        if x == 1 {
            next = 2
        } else {
            next = 1
        }
        for _, j := range dislikeMap[n] {
            dfs(j, next)
        }
    }

    for i := 1; i <= n; i++ {
        if visited[i] == 0 {
            x := 0
            if len(groups[0]) <= len(groups[1]) {
                x = 1
            } else {
                x = 2
            }
            dfs(i, x)
        }

    }

    fmt.Println(groups)
    return len(groups[0]) == len(groups[1])
}
