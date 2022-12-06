package leetcode


/**
1,3,-1,-3,5,3,6,7

单调队列的思想，核心是队列存储按照大小排序的位置index，而不是实际的数

 */

func MaxSlidingWindow(nums []int, k int) []int {
    q := []int{}

    push := func(i int) {
        for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }

    //确定第一个窗口的队列
    for i := 0; i < k; i++ {
        push(i)
    }

    res := make([]int, len(nums)-k+1)
    res[0] = nums[q[0]]
    for i := k; i < len(nums); i++ {
        if q[0] <= i-k {
            q = q[1:]
        }
        push(i)
        res[i-k+1] = nums[q[0]]
    }
    return res
}

func GF_MaxSlidingWindow(nums []int, k int) []int {
    q := []int{}
    push := func(i int) {
        for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }

    for i := 0; i < k; i++ {
        push(i)
    }

    n := len(nums)
    ans := make([]int, 1, n-k+1)
    ans[0] = nums[q[0]]
    for i := k; i < n; i++ {
        push(i)
        for q[0] <= i-k {
            q = q[1:]
        }
        ans = append(ans, nums[q[0]])
    }
    return ans
}

