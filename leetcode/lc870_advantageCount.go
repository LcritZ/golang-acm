package leetcode

import "sort"

/**
给定两个大小相等的数组A和B，A 相对于 B 的优势可以用满足A[i] > B[i]的索引 i的数目来描述。

返回A的任意排列，使其相对于 B的优势最大化。

对A排序，然后按照大小， 最小的给大于当前不能满足B的

排序+二分

A排序，然后遍历B，每个数据找到A的最相近一个， 用完一个A的数就更新排序后A数组，不能重复用一个大数比较


输入：A = [2,7,11,15], B = [1,10,4,11]
输出：[2,11,7,15]

 */
//func advantageCount(nums1 []int, nums2 []int) []int {
//
//}

func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	res := make([]int, len(nums2))
	for id ,num := range nums2 {
		target := searchInt(nums1, num)
		if target < len(nums1) {
			res[id] = nums1[target]
			nums1 = append(nums1[0:target], nums1[target+1:]...)
		} else {
			res[id] = nums1[0]
			nums1 = nums1[1:]
		}
	}
	return res
}

func searchInt(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] {
		return 0
	}
	if target >= nums[len(nums) - 1] {
		return len(nums)
	}

	left, right := 0, len(nums) - 1
	for right - left > 1 {
		mid := (left + right + 1)/2
		if nums[mid] > target {
			right = mid
		}
		if nums[mid] <= target {
			left = mid
		}
	}
	return right
}
