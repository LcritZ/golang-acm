package leetcode

/**
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素

 */
func RemoveDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slow, fast := 0, 0
    for fast < len(nums)  {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
        fast++
    }

    return slow+1
}

