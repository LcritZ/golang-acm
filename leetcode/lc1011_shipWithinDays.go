package leetcode

import "fmt"

/**
输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。

输入：weights = [3,2,2,4,1,4], days = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4

思路：二分的方法，用函数判断是否满足条件，左右区间再处理，左闭右开， right需要=mid

确定left=1， right最大值就等于 length/days * weights中最大值 也就是一天就带走几次的重量，而且每次都是最大值，就能满足其他的都带走

条件给出了weight最大是500 
 */
func ShipWithinDays(weights []int, days int) int {
    max := 500*((len(weights)-1)/days + 1)
    min := 1

    for min < max {
        mid := min + (max-min)/2
        fmt.Println("mid:", mid, "min:", min, "max:", max)
        if ShipHelper(weights, days, mid) {
            max = mid
        } else {
            min = mid+1
        }
    }
    return min
}

func ShipHelper(weights []int, days int, loads int) bool {
    sum := weights[0]
    costDays := 0
    for i := 1; i < len(weights); i++ {
        if sum > loads {
            return false
        }
        if (sum+weights[i]) > loads  {
            costDays++
            sum = 0
            fmt.Println("i:", i, "cost:", costDays)
        }
        sum += weights[i]
    }
    if sum > loads {
        return false
    }
    if sum > 0 {
        costDays++
    }
    return costDays <= days
}
