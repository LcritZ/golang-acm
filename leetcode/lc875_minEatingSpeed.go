package leetcode

import "math"

/**
珂珂喜欢吃香蕉。这里有N堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在H小时后回来。

珂珂可以决定她吃香蕉的速度K（单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）

输入: piles = [3,6,7,11], H = 8
输出: 4
全部加起来 27 除以8 = 4

1 2 2 3

2 3 3

1 1 1

1 <= piles.length <= 10^4
piles.length <= H <= 10^9
1 <= piles[i] <= 10^9

 */
func MinEatingSpeed(piles []int, h int) int {
    min, max := 1, int(math.Pow10(9))
    for min < max {
        mid := min + (max-min)/2
        //直接判断改为用函数判断
        if Eatinghelper(piles, h, mid) {
            max = mid
        } else {
            min = mid+1
        }

    }
    return min
}


//将计算用函数标识
func Eatinghelper(piles []int, h int, k int) bool {
    times := 0
    for i := 0; i < len(piles); i++ {
        //技巧：先-1，这时候除就避免了整除的时候多加1的情况
        times += (piles[i]-1)/k + 1
    }
    return times <= h
}
