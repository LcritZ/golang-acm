package basic

/**

1-100个 递增数组

32位之内，

100个列表内找第K个最大的数

思路：
构造最大堆，100个最后的数加入堆

K个最大，K/100 = x  每组最后x个

[1,2,3]
[4,5,6]

log100*log100 *k

快排+二分找到

 */

/**
1,2,3,5

整数n, 小于n最大素数

 */

func GetMax(n int) int {
    if n == 1 {
        return 0
    }

    if n <= 3 {
        return n-1
    }

    nums := []int{2,3}
    for i := 4; i < n; i++ {
        flag := true
        for _, x := range nums {
            if i % x == 0  {
                flag = false
                break
            }
        }
        if flag {
            nums = append(nums, i)
        }
    }

    return nums[len(nums)-1]

}
