package leetcode

import "golang-acm/util"

/**
使用如下判断只能解决部分问题
-2,3,-4这种就返回3，但实际最大是24 -6*-4=24

 */
//func MaxProduct(nums []int) int {
//    if len(nums) == 0 {
//        return 0
//    }
//
//    curr := nums[0]
//    res := curr
//    for i := 1; i < len(nums); i++ {
//        fmt.Println(curr)
//        if nums[i]*curr > 0 {
//            if nums[i]*curr >= curr {
//                curr = nums[i]*curr
//            }
//        } else {
//            curr = nums[i]
//        }
//        res = util.Max(res, curr)
//    }
//
//    return res
//}

func MaxProduct(nums []int) int {
    max, min, res := nums[0], nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        x, y := max, min
        max = util.Max(nums[i]*max, util.Max(nums[i], nums[i]*y))
        min = util.Min(nums[i]*min, util.Min(nums[i], nums[i]*x))
        res = util.Max(res, max)
    }

    return res

}