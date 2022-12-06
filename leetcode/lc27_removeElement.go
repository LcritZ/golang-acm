package leetcode

/**
原地删除数组元素
通过交换的方法，将指定的元素放在最后面
 */
func RemoveElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    slow, fast := 0, 0
    for fast < len(nums)  {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }

    return slow
}

