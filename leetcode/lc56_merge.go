package leetcode

import (
    "golang-acm/util"
    "sort"
)

/**
2端比较大小，就先排序确定一端，再处理另一端

 */

func merge56(intervals [][]int) [][]int {
    //先从小到大排序
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0]<intervals[j][0]
    })
    //再弄重复的
    for i:=0;i<len(intervals)-1;i++{
        if intervals[i][1]>=intervals[i+1][0]{
            intervals[i][1]=util.Max(intervals[i][1],intervals[i+1][1])//赋值最大值
            intervals=append(intervals[:i+1],intervals[i+2:]...)
            i--
        }
    }
    return intervals
}

func Merge56_2(intervals [][]int) [][]int  {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    for i := 0; i < len(intervals)-1; {
        if intervals[i][1] >= intervals[i+1][0] {
            intervals[i][1] = util.Max(intervals[i][1], intervals[i+1][1])
            intervals = append(intervals[:i+1],intervals[i+2:]...)
            //符合比较，i不变
        } else {
            i++
        }
    }

    return intervals

}
