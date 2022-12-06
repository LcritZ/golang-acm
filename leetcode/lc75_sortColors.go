package leetcode

/**
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库的sort函数的情况下解决这个问题
输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

思路：双指针，走2次遍历，实现0,1替换
 */

func SortColors(nums []int)  {
    if len(nums) == 0 {
        return
    }

    left, right := 0, 0
    for ; right < len(nums); right++ {
        if nums[right] == 0 {
            temp := nums[left]
            nums[left] = 0
            nums[right] = temp
            left++
        }
    }

    right = left
    for ; right < len(nums); right++ {
        if nums[right] == 1 {
            temp := nums[left]
            nums[left] = 1
            nums[right] = temp
            left++
        }
    }

}
