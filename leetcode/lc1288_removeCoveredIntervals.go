package leetcode

import (
    "sort"
)

/**
给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。

输入：intervals = [[1,4],[3,6],[2,8]]
输出：2
解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。

这种区间合并，变更的问题需要注意修改遍历的I，因为默认++，但是数组变了，需要重新遍历

 */
func RemoveCoveredIntervals(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    for i := 0; i < len(intervals)-1; i++ {
        if intervals[i+1][0] >= intervals[i][0] && intervals[i+1][1] <= intervals[i][1]   {
            intervals = append(intervals[:i+1], intervals[i+2:]...)
            i--
            continue
        }
        if intervals[i+1][0] == intervals[i][0] && intervals[i+1][1] >= intervals[i][1] {
            intervals = append(intervals[:i], intervals[i+1:]...)
            i--
            continue
        }
    }

    return len(intervals)

}

